package provider

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/config"
	"github.com/itsLeonB/orcashtrator/internal/domain/auth"
	"github.com/itsLeonB/orcashtrator/internal/domain/debt"
	"github.com/itsLeonB/orcashtrator/internal/domain/expensebill"
	"github.com/itsLeonB/orcashtrator/internal/domain/expenseitem"
	"github.com/itsLeonB/orcashtrator/internal/domain/friendship"
	"github.com/itsLeonB/orcashtrator/internal/domain/groupexpense"
	"github.com/itsLeonB/orcashtrator/internal/domain/imageupload"
	"github.com/itsLeonB/orcashtrator/internal/domain/otherfee"
	"github.com/itsLeonB/orcashtrator/internal/domain/profile"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type Clients struct {
	conns          []*grpc.ClientConn
	Auth           auth.AuthClient
	Profile        profile.ProfileClient
	Friendship     friendship.FriendshipClient
	TransferMethod debt.TransferMethodClient
	Debt           debt.DebtClient
	GroupExpense   groupexpense.GroupExpenseClient
	ExpenseItem    expenseitem.ExpenseItemClient
	OtherFee       otherfee.OtherFeeClient
	ExpenseBill    expensebill.ExpenseBillClient
	ImageUpload    imageupload.ImageUploadClient
}

func ProvideClients(configs config.ServiceClient, validate *validator.Validate, logger ezutil.Logger) *Clients {
	conns := make([]*grpc.ClientConn, 0, 3)

	billsplittrConn, err := grpc.NewClient(
		configs.BillsplittrHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Warnf("error connecting to billsplittr client: %v", err)
	} else {
		conns = append(conns, billsplittrConn)
	}

	cocoonConn, err := grpc.NewClient(
		configs.CocoonHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Warnf("error connecting to cocoon client: %v", err)
	} else {
		conns = append(conns, cocoonConn)
	}

	drexConn, err := grpc.NewClient(
		configs.DrexHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Warnf("error connecting to drex client: %v", err)
	} else {
		conns = append(conns, drexConn)
	}

	stortrConn, err := grpc.NewClient(
		configs.StortrHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             5 * time.Second,
			PermitWithoutStream: false,
		}),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(appconstant.MaxFileSize),
			grpc.MaxCallSendMsgSize(appconstant.MaxFileSize),
		),
	)
	if err != nil {
		logger.Warnf("error connecting to stortr client: %v", err)
	} else {
		conns = append(conns, stortrConn)
	}

	return &Clients{
		conns,
		auth.NewAuthClient(validate, cocoonConn),
		profile.NewProfileClient(validate, cocoonConn),
		friendship.NewFriendshipClient(validate, cocoonConn),
		debt.NewTransferMethodClient(drexConn),
		debt.NewDebtClient(validate, drexConn),
		groupexpense.NewGroupExpenseClient(validate, billsplittrConn),
		expenseitem.NewExpenseItemClient(validate, billsplittrConn),
		otherfee.NewOtherFeeClient(validate, billsplittrConn),
		expensebill.NewExpenseBillClient(validate, billsplittrConn),
		imageupload.NewImageUploadClient(validate, stortrConn),
	}
}

func (c *Clients) Shutdown() error {
	var err error
	for _, conn := range c.conns {
		if e := conn.Close(); e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}

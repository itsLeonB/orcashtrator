package debt

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/itsLeonB/drex-protos/gen/go/debt/v1"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/rotisserie/eris"
	"golang.org/x/text/currency"
	"google.golang.org/grpc"
)

type DebtClient interface {
	RecordNewTransaction(ctx context.Context, req RecordNewTransactionRequest) (Transaction, error)
	GetTransactions(ctx context.Context, profileID uuid.UUID) ([]Transaction, error)
	ProcessConfirmedGroupExpense(ctx context.Context, groupExpense GroupExpenseData) error
	GetAllByProfileIDs(ctx context.Context, req GetAllByProfileIDsRequest) ([]Transaction, error)
}

type debtClient struct {
	validate *validator.Validate
	client   debt.DebtServiceClient
}

func NewDebtClient(validate *validator.Validate, conn *grpc.ClientConn) DebtClient {
	if validate == nil {
		panic("validator is nil")
	}

	return &debtClient{
		validate,
		debt.NewDebtServiceClient(conn),
	}
}

func (dc *debtClient) RecordNewTransaction(ctx context.Context, req RecordNewTransactionRequest) (Transaction, error) {
	if err := dc.validate.Struct(req); err != nil {
		return Transaction{}, err
	}

	action, err := toTransactionActionEnum(req.Action)
	if err != nil {
		return Transaction{}, err
	}

	request := &debt.RecordNewTransactionRequest{
		UserProfileId:    req.UserProfileID.String(),
		FriendProfileId:  req.FriendProfileID.String(),
		Action:           action,
		Amount:           ezutil.DecimalToMoney(req.Amount, currency.IDR.String()),
		TransferMethodId: req.TransferMethodID.String(),
		Description:      req.Description,
	}

	response, err := dc.client.RecordNewTransaction(ctx, request)
	if err != nil {
		return Transaction{}, err
	}

	return fromTransactionProto(response.GetTransaction())
}

func (dc *debtClient) GetTransactions(ctx context.Context, profileID uuid.UUID) ([]Transaction, error) {
	if profileID == uuid.Nil {
		return nil, eris.New("profile id is nil")
	}

	request := &debt.GetTransactionsRequest{
		UserProfileId: profileID.String(),
	}

	response, err := dc.client.GetTransactions(ctx, request)
	if err != nil {
		return nil, err
	}

	return ezutil.MapSliceWithError(response.GetTransactions(), fromTransactionProto)
}

func (dc *debtClient) ProcessConfirmedGroupExpense(ctx context.Context, groupExpense GroupExpenseData) error {
	if err := dc.validate.Struct(groupExpense); err != nil {
		return err
	}

	request := &debt.ProcessConfirmedGroupExpenseRequest{
		GroupExpense: &debt.GroupExpenseData{
			Id:               groupExpense.ID.String(),
			PayerProfileId:   groupExpense.PayerProfileID.String(),
			CreatorProfileId: groupExpense.CreatorProfileID.String(),
			Description:      groupExpense.Description,
			Participants:     ezutil.MapSlice(groupExpense.Participants, toExpenseParticipantProto),
		},
	}

	_, err := dc.client.ProcessConfirmedGroupExpense(ctx, request)

	return err
}

func (dc *debtClient) GetAllByProfileIDs(ctx context.Context, req GetAllByProfileIDsRequest) ([]Transaction, error) {
	if err := dc.validate.Struct(req); err != nil {
		return nil, err
	}

	request := &debt.GetAllByProfileIdsRequest{
		UserProfileId:   req.UserProfileID.String(),
		FriendProfileId: req.FriendProfileID.String(),
	}

	response, err := dc.client.GetAllByProfileIds(ctx, request)
	if err != nil {
		return nil, err
	}

	return ezutil.MapSliceWithError(response.GetTransactions(), fromTransactionProto)
}

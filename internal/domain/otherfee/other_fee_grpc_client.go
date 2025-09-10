package otherfee

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/itsLeonB/billsplittr-protos/gen/go/otherfee/v1"
	"github.com/itsLeonB/ezutil/v2"
	"google.golang.org/grpc"
)

type OtherFeeClient interface {
	Add(ctx context.Context, req AddRequest) (OtherFee, error)
	Update(ctx context.Context, req UpdateRequest) (OtherFee, error)
	Remove(ctx context.Context, req RemoveRequest) error
	GetCalculationMethods(ctx context.Context) ([]CalculationMethodInfo, error)
}

type otherFeeClient struct {
	validate *validator.Validate
	client   otherfee.OtherFeeServiceClient
}

func NewOtherFeeClient(
	validate *validator.Validate,
	conn *grpc.ClientConn,
) OtherFeeClient {
	if validate == nil {
		panic("validator is nil")
	}

	return &otherFeeClient{
		validate,
		otherfee.NewOtherFeeServiceClient(conn),
	}
}

func (ofc *otherFeeClient) Add(ctx context.Context, req AddRequest) (OtherFee, error) {
	if err := ofc.validate.Struct(req); err != nil {
		return OtherFee{}, err
	}

	otherFeeData, err := ToOtherFeeProto(req.OtherFeeData)
	if err != nil {
		return OtherFee{}, err
	}

	request := &otherfee.AddRequest{
		ProfileId:      req.ProfileID.String(),
		GroupExpenseId: req.GroupExpenseID.String(),
		OtherFee:       otherFeeData,
	}

	response, err := ofc.client.Add(ctx, request)
	if err != nil {
		return OtherFee{}, err
	}

	return FromOtherFeeResponseProto(response.GetOtherFee())
}

func (ofc *otherFeeClient) Update(ctx context.Context, req UpdateRequest) (OtherFee, error) {
	if err := ofc.validate.Struct(req); err != nil {
		return OtherFee{}, err
	}

	otherFeeData, err := ToOtherFeeProto(req.OtherFeeData)
	if err != nil {
		return OtherFee{}, err
	}

	request := &otherfee.UpdateRequest{
		ProfileId:      req.ProfileID.String(),
		Id:             req.ID.String(),
		GroupExpenseId: req.GroupExpenseID.String(),
		OtherFee:       otherFeeData,
	}

	response, err := ofc.client.Update(ctx, request)
	if err != nil {
		return OtherFee{}, err
	}

	return FromOtherFeeResponseProto(response.GetOtherFee())
}

func (ofc *otherFeeClient) Remove(ctx context.Context, req RemoveRequest) error {
	if err := ofc.validate.Struct(req); err != nil {
		return err
	}

	request := &otherfee.RemoveRequest{
		ProfileId:      req.ProfileID.String(),
		Id:             req.ID.String(),
		GroupExpenseId: req.GroupExpenseID.String(),
	}

	_, err := ofc.client.Remove(ctx, request)

	return err
}

func (ofc *otherFeeClient) GetCalculationMethods(ctx context.Context) ([]CalculationMethodInfo, error) {
	response, err := ofc.client.GetCalculationMethods(ctx, nil)
	if err != nil {
		return nil, err
	}

	return ezutil.MapSliceWithError(response.GetMethods(), fromCalculationMethodInfoProto)
}

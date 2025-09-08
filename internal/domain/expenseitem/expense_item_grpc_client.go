package expenseitem

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/itsLeonB/billsplittr-protos/gen/go/expenseitem/v1"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/rotisserie/eris"
	"google.golang.org/grpc"
)

type ExpenseItemClient interface {
	Add(ctx context.Context, req AddRequest) (ExpenseItem, error)
	GetDetails(ctx context.Context, req GetDetailsRequest) (ExpenseItem, error)
	Update(ctx context.Context, req UpdateRequest) (ExpenseItem, error)
	Remove(ctx context.Context, req RemoveRequest) error
}

type expenseItemClient struct {
	validate *validator.Validate
	client   expenseitem.ExpenseItemServiceClient
}

func NewExpenseItemClient(
	validate *validator.Validate,
	conn *grpc.ClientConn,
) ExpenseItemClient {
	return &expenseItemClient{
		validate,
		expenseitem.NewExpenseItemServiceClient(conn),
	}
}

func (eic *expenseItemClient) Add(ctx context.Context, req AddRequest) (ExpenseItem, error) {
	if err := eic.validate.Struct(req); err != nil {
		return ExpenseItem{}, err
	}

	request := &expenseitem.AddRequest{
		ProfileId:      req.ProfileID.String(),
		GroupExpenseId: req.GroupExpenseID.String(),
		ExpenseItem:    ToExpenseItemProto(req.ExpenseItemData),
	}

	response, err := eic.client.Add(ctx, request)
	if err != nil {
		return ExpenseItem{}, err
	}

	return FromExpenseItemResponseProto(response.GetExpenseItem())
}

func (eic *expenseItemClient) GetDetails(ctx context.Context, req GetDetailsRequest) (ExpenseItem, error) {
	if err := eic.validate.Struct(req); err != nil {
		return ExpenseItem{}, err
	}

	request := &expenseitem.GetDetailsRequest{
		Id:             req.ID.String(),
		GroupExpenseId: req.GroupExpenseID.String(),
	}

	response, err := eic.client.GetDetails(ctx, request)
	if err != nil {
		return ExpenseItem{}, err
	}

	return FromExpenseItemResponseProto(response.GetExpenseItem())
}

func (eic *expenseItemClient) Update(ctx context.Context, req UpdateRequest) (ExpenseItem, error) {
	if err := eic.validate.Struct(req); err != nil {
		return ExpenseItem{}, err
	}

	request := &expenseitem.UpdateRequest{
		ProfileId:      req.ProfileID.String(),
		Id:             req.ID.String(),
		GroupExpenseId: req.GroupExpenseID.String(),
		ExpenseItem:    ToExpenseItemProto(req.ExpenseItemData),
	}

	response, err := eic.client.Update(ctx, request)
	if err != nil {
		return ExpenseItem{}, eris.Wrap(err, appconstant.ErrServiceClient)
	}

	return FromExpenseItemResponseProto(response.GetExpenseItem())
}

func (eic *expenseItemClient) Remove(ctx context.Context, req RemoveRequest) error {
	if err := eic.validate.Struct(req); err != nil {
		return err
	}

	request := &expenseitem.RemoveRequest{
		ProfileId:      req.ProfileID.String(),
		Id:             req.ID.String(),
		GroupExpenseId: req.GroupExpenseID.String(),
	}

	_, err := eic.client.Remove(ctx, request)

	return err
}

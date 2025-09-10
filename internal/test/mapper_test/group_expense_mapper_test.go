package mapper_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGroupExpenseRequestToEntity(t *testing.T) {
	request := dto.NewGroupExpenseRequest{
		PayerProfileID: uuid.New(),
		TotalAmount:    decimal.NewFromInt(1000),
		Subtotal:       decimal.NewFromInt(900),
		Description:    "Test expense",
		Items: []dto.NewExpenseItemRequest{
			{
				Name:     "Item 1",
				Amount:   decimal.NewFromInt(500),
				Quantity: 1,
			},
		},
		OtherFees: []dto.NewOtherFeeRequest{
			{
				Name:   "Service Fee",
				Amount: decimal.NewFromInt(100),
			},
		},
	}

	result := mapper.GroupExpenseRequestToEntity(request)

	assert.Equal(t, request.PayerProfileID, result.PayerProfileID)
	assert.Equal(t, request.TotalAmount, result.TotalAmount)
	assert.Equal(t, request.Subtotal, result.Subtotal)
	assert.Equal(t, request.Description, result.Description)
	assert.Len(t, result.Items, 1)
	assert.Equal(t, request.Items[0].Name, result.Items[0].Name)
	assert.Equal(t, request.Items[0].Amount, result.Items[0].Amount)
	assert.Equal(t, request.Items[0].Quantity, result.Items[0].Quantity)
	assert.Len(t, result.OtherFees, 1)
}

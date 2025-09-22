package mapper

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/domain/expensebill"
	"github.com/itsLeonB/orcashtrator/internal/dto"
)

func ExpenseBillToResponse(bill expensebill.ExpenseBill, url string, profileID uuid.UUID) dto.ExpenseBillResponse {
	return dto.ExpenseBillResponse{
		ID:               bill.ID,
		CreatorProfileID: bill.CreatorProfileID,
		PayerProfileID:   bill.PayerProfileID,
		ImageURL:         url,
		CreatedAt:        bill.CreatedAt,
		UpdatedAt:        bill.UpdatedAt,
		DeletedAt:        bill.DeletedAt,
		IsCreatedByUser:  bill.CreatorProfileID == profileID,
		IsPaidByUser:     bill.PayerProfileID == profileID,
	}
}

func ExpenseBillSimpleMapper(url string, profileID uuid.UUID) func(expensebill.ExpenseBill) dto.ExpenseBillResponse {
	return func(eb expensebill.ExpenseBill) dto.ExpenseBillResponse {
		return ExpenseBillToResponse(eb, url, profileID)
	}
}

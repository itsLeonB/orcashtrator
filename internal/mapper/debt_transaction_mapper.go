package mapper

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/domain/debt"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/shopspring/decimal"
	"golang.org/x/text/currency"
)

func MapToFriendBalanceSummary(userProfileID uuid.UUID, debtTransactions []dto.DebtTransactionResponse) dto.FriendBalance {
	totalOwedToYou, totalYouOwe := decimal.Zero, decimal.Zero

	for _, transaction := range debtTransactions {
		switch transaction.Type {
		case appconstant.Lend:
			switch transaction.Action {
			case appconstant.LendAction: // You lent money
				totalOwedToYou = totalOwedToYou.Add(transaction.Amount)
			case appconstant.BorrowAction: // You borrowed money
				totalYouOwe = totalYouOwe.Add(transaction.Amount)
			}
		case appconstant.Repay:
			switch transaction.Action {
			case appconstant.ReceiveAction: // You received repayment
				totalOwedToYou = totalOwedToYou.Sub(transaction.Amount)
			case appconstant.ReturnAction: // You returned money
				totalYouOwe = totalYouOwe.Sub(transaction.Amount)
			}
		}
	}

	return dto.FriendBalance{
		TotalOwedToYou: totalOwedToYou,
		TotalYouOwe:    totalYouOwe,
		NetBalance:     totalOwedToYou.Sub(totalYouOwe),
		CurrencyCode:   currency.IDR.String(),
	}
}

func DebtTransactionToResponse(dt debt.Transaction) dto.DebtTransactionResponse {
	return dto.DebtTransactionResponse{
		ID:             dt.ID,
		ProfileID:      dt.ProfileID,
		Type:           dt.Type,
		Action:         dt.Action,
		Amount:         dt.Amount,
		TransferMethod: dt.TransferMethod,
		Description:    dt.Description,
		CreatedAt:      dt.CreatedAt,
		UpdatedAt:      dt.UpdatedAt,
		DeletedAt:      dt.DeletedAt,
	}
}

package expenseitem

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/billsplittr-protos/gen/go/expenseitem/v1"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/domain"
	"github.com/rotisserie/eris"
	"github.com/shopspring/decimal"
	"golang.org/x/text/currency"
)

func ToExpenseItemProto(ei ExpenseItemData) *expenseitem.ExpenseItem {
	return &expenseitem.ExpenseItem{
		Name:         ei.Name,
		Amount:       ezutil.DecimalToMoney(ei.Amount, currency.IDR.String()),
		Quantity:     int64(ei.Quantity),
		Participants: ezutil.MapSlice(ei.Participants, toItemParticipantProto),
	}
}

func FromExpenseItemResponseProto(eir *expenseitem.ExpenseItemResponse) (ExpenseItem, error) {
	if eir == nil {
		return ExpenseItem{}, eris.New("expense item response is nil")
	}

	groupExpenseID, err := ezutil.Parse[uuid.UUID](eir.GetGroupExpenseId())
	if err != nil {
		return ExpenseItem{}, err
	}

	expenseItem := eir.GetExpenseItem()
	if expenseItem == nil {
		return ExpenseItem{}, eris.New("expense item data is nil")
	}

	participants, err := ezutil.MapSliceWithError(expenseItem.GetParticipants(), fromItemParticipantProto)
	if err != nil {
		return ExpenseItem{}, err
	}

	metadata, err := domain.FromAuditMetadataProto(eir.GetAuditMetadata())
	if err != nil {
		return ExpenseItem{}, err
	}

	return ExpenseItem{
		GroupExpenseID: groupExpenseID,
		ExpenseItemData: ExpenseItemData{
			Name:         expenseItem.GetName(),
			Amount:       ezutil.MoneyToDecimal(expenseItem.GetAmount()),
			Quantity:     int(expenseItem.GetQuantity()),
			Participants: participants,
		},
		AuditMetadata: metadata,
	}, nil
}

func fromItemParticipantProto(ipr *expenseitem.ItemParticipant) (ItemParticipant, error) {
	if ipr == nil {
		return ItemParticipant{}, eris.New("item participant is nil")
	}

	profileID, err := ezutil.Parse[uuid.UUID](ipr.GetProfileId())
	if err != nil {
		return ItemParticipant{}, err
	}

	return ItemParticipant{
		ProfileID: profileID,
		Share:     decimal.NewFromFloat(ipr.GetShare()),
	}, nil
}

func toItemParticipantProto(ip ItemParticipant) *expenseitem.ItemParticipant {
	return &expenseitem.ItemParticipant{
		ProfileId: ip.ProfileID.String(),
		Share:     ip.Share.InexactFloat64(),
	}
}

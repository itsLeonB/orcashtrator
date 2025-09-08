package mapper

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/domain/debt"
	"github.com/itsLeonB/orcashtrator/internal/domain/expenseitem"
	"github.com/itsLeonB/orcashtrator/internal/domain/groupexpense"
	"github.com/itsLeonB/orcashtrator/internal/domain/otherfee"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/ungerr"
	"github.com/shopspring/decimal"
)

func GroupExpenseRequestToEntity(request dto.NewGroupExpenseRequest) groupexpense.CreateDraftRequest {
	return groupexpense.CreateDraftRequest{
		PayerProfileID: request.PayerProfileID,
		TotalAmount:    request.TotalAmount,
		Subtotal:       request.Subtotal,
		Description:    request.Description,
		Items:          ezutil.MapSlice(request.Items, NewExpenseItemRequestToData),
		OtherFees:      ezutil.MapSlice(request.OtherFees, otherFeeRequestToData),
	}
}

func GroupExpenseToResponse(groupExpense groupexpense.GroupExpense, userProfileID uuid.UUID, namesByProfileID map[uuid.UUID]string) dto.GroupExpenseResponse {
	return dto.GroupExpenseResponse{
		ID:                    groupExpense.ID,
		PayerProfileID:        groupExpense.PayerProfileID,
		PayerName:             namesByProfileID[groupExpense.PayerProfileID],
		PaidByUser:            groupExpense.PayerProfileID == userProfileID,
		TotalAmount:           groupExpense.TotalAmount,
		Description:           groupExpense.Description,
		Items:                 ezutil.MapSlice(groupExpense.Items, getExpenseItemSimpleMapper(userProfileID, namesByProfileID)),
		OtherFees:             ezutil.MapSlice(groupExpense.OtherFees, getOtherFeeSimpleMapper(userProfileID, namesByProfileID)),
		CreatorProfileID:      groupExpense.CreatorProfileID,
		CreatorName:           namesByProfileID[groupExpense.CreatorProfileID],
		CreatedByUser:         groupExpense.CreatorProfileID == userProfileID,
		Confirmed:             groupExpense.IsConfirmed,
		ParticipantsConfirmed: groupExpense.IsParticipantsConfirmed,
		CreatedAt:             groupExpense.CreatedAt,
		UpdatedAt:             groupExpense.UpdatedAt,
		DeletedAt:             groupExpense.DeletedAt,
		Participants:          ezutil.MapSlice(groupExpense.Participants, getExpenseParticipantSimpleMapper(userProfileID, namesByProfileID)),
	}
}

func getExpenseItemSimpleMapper(userProfileID uuid.UUID, namesByProfileID map[uuid.UUID]string) func(item expenseitem.ExpenseItem) dto.ExpenseItemResponse {
	return func(item expenseitem.ExpenseItem) dto.ExpenseItemResponse {
		return ExpenseItemToResponse(item, userProfileID, namesByProfileID)
	}
}

func ExpenseItemToResponse(item expenseitem.ExpenseItem, userProfileID uuid.UUID, namesByProfileID map[uuid.UUID]string) dto.ExpenseItemResponse {
	return dto.ExpenseItemResponse{
		ID:             item.ID,
		GroupExpenseID: item.GroupExpenseID,
		Name:           item.Name,
		Amount:         item.Amount,
		Quantity:       item.Quantity,
		CreatedAt:      item.CreatedAt,
		UpdatedAt:      item.UpdatedAt,
		DeletedAt:      item.DeletedAt,
		Participants:   ezutil.MapSlice(item.Participants, getItemParticipantSimpleMapper(userProfileID, namesByProfileID)),
	}
}

func getOtherFeeSimpleMapper(userProfileID uuid.UUID, namesByProfileID map[uuid.UUID]string) func(otherfee.OtherFee) dto.OtherFeeResponse {
	return func(fee otherfee.OtherFee) dto.OtherFeeResponse {
		return OtherFeeToResponse(fee, userProfileID, namesByProfileID)
	}
}

func OtherFeeToResponse(fee otherfee.OtherFee, userProfileID uuid.UUID, namesByProfileID map[uuid.UUID]string) dto.OtherFeeResponse {
	return dto.OtherFeeResponse{
		ID:                fee.ID,
		Name:              fee.Name,
		Amount:            fee.Amount,
		CalculationMethod: fee.CalculationMethod,
		CreatedAt:         fee.CreatedAt,
		UpdatedAt:         fee.UpdatedAt,
		DeletedAt:         fee.DeletedAt,
		Participants:      ezutil.MapSlice(fee.Participants, getFeeParticipantSimpleMapper(userProfileID, namesByProfileID)),
	}
}

func getFeeParticipantSimpleMapper(userProfileID uuid.UUID, namesByProfileID map[uuid.UUID]string) func(otherfee.FeeParticipant) dto.FeeParticipantResponse {
	return func(feeParticipant otherfee.FeeParticipant) dto.FeeParticipantResponse {
		return feeParticipantToResponse(feeParticipant, userProfileID, namesByProfileID[feeParticipant.ProfileID])
	}
}

func feeParticipantToResponse(feeParticipant otherfee.FeeParticipant, userProfileID uuid.UUID, participantProfileName string) dto.FeeParticipantResponse {
	return dto.FeeParticipantResponse{
		ProfileName: participantProfileName,
		ProfileID:   feeParticipant.ProfileID,
		ShareAmount: feeParticipant.ShareAmount,
		IsUser:      feeParticipant.ProfileID == userProfileID,
	}
}

func getItemParticipantSimpleMapper(userProfileID uuid.UUID, namesByProfileID map[uuid.UUID]string) func(itemParticipant expenseitem.ItemParticipant) dto.ItemParticipantResponse {
	return func(itemParticipant expenseitem.ItemParticipant) dto.ItemParticipantResponse {
		return itemParticipantToResponse(itemParticipant, userProfileID, namesByProfileID[itemParticipant.ProfileID])
	}
}

func itemParticipantToResponse(itemParticipant expenseitem.ItemParticipant, userProfileID uuid.UUID, participantProfileName string) dto.ItemParticipantResponse {
	return dto.ItemParticipantResponse{
		ProfileName: participantProfileName,
		ProfileID:   itemParticipant.ProfileID,
		Share:       itemParticipant.Share,
		IsUser:      itemParticipant.ProfileID == userProfileID,
	}
}

func otherFeeRequestToData(req dto.NewOtherFeeRequest) otherfee.OtherFeeData {
	return otherfee.OtherFeeData{
		Name:              req.Name,
		Amount:            req.Amount,
		CalculationMethod: req.CalculationMethod,
	}
}

func OtherFeeRequestToEntity(request dto.NewOtherFeeRequest) otherfee.OtherFee {
	return otherfee.OtherFee{
		GroupExpenseID: request.GroupExpenseID,
		OtherFeeData:   otherFeeRequestToData(request),
	}
}

func PatchExpenseItemWithRequest(expenseItem expenseitem.ExpenseItem, request dto.UpdateExpenseItemRequest) expenseitem.ExpenseItem {
	expenseItem.Name = request.Name
	expenseItem.Amount = request.Amount
	expenseItem.Quantity = request.Quantity
	return expenseItem
}

func ItemParticipantRequestToEntity(itemParticipant dto.ItemParticipantRequest) expenseitem.ItemParticipant {
	return expenseitem.ItemParticipant{
		ProfileID: itemParticipant.ProfileID,
		Share:     itemParticipant.Share,
	}
}

func ExpenseParticipantToResponse(expenseParticipant groupexpense.ExpenseParticipant, userProfileID uuid.UUID, participantProfileName string) dto.ExpenseParticipantResponse {
	return dto.ExpenseParticipantResponse{
		ProfileName: participantProfileName,
		ProfileID:   expenseParticipant.ProfileID,
		ShareAmount: expenseParticipant.ShareAmount,
		IsUser:      expenseParticipant.ProfileID == userProfileID,
	}
}

func getExpenseParticipantSimpleMapper(userProfileID uuid.UUID, namesByProfileID map[uuid.UUID]string) func(groupexpense.ExpenseParticipant) dto.ExpenseParticipantResponse {
	return func(ep groupexpense.ExpenseParticipant) dto.ExpenseParticipantResponse {
		return ExpenseParticipantToResponse(ep, userProfileID, namesByProfileID[ep.ProfileID])
	}
}

func PatchOtherFeeWithRequest(otherFee otherfee.OtherFee, request dto.UpdateOtherFeeRequest) otherfee.OtherFee {
	otherFee.Name = request.Name
	otherFee.Amount = request.Amount
	otherFee.CalculationMethod = request.CalculationMethod
	return otherFee
}

func ExpenseParticipantToData(participant groupexpense.ExpenseParticipant) (debt.ExpenseParticipantData, error) {
	if participant.ShareAmount.LessThanOrEqual(decimal.Zero) {
		return debt.ExpenseParticipantData{}, ungerr.UnprocessableEntityError(fmt.Sprintf(
			"participant %s has share amount: %s",
			participant.ProfileID,
			participant.ShareAmount.String(),
		))
	}
	return debt.ExpenseParticipantData{
		ProfileID:   participant.ProfileID,
		ShareAmount: participant.ShareAmount,
	}, nil
}

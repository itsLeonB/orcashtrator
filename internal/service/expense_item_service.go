package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/domain/expenseitem"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
	"github.com/itsLeonB/orcashtrator/internal/util"
	"github.com/itsLeonB/ungerr"
)

type expenseItemServiceImpl struct {
	profileService    ProfileService
	expenseItemClient expenseitem.ExpenseItemClient
}

func NewExpenseItemService(
	profileService ProfileService,
	expenseItemClient expenseitem.ExpenseItemClient,
) ExpenseItemService {
	return &expenseItemServiceImpl{
		profileService,
		expenseItemClient,
	}
}

func (ges *expenseItemServiceImpl) Add(ctx context.Context, req dto.NewExpenseItemRequest) (dto.ExpenseItemResponse, error) {
	profileID, err := util.GetProfileID(ctx)
	if err != nil {
		return dto.ExpenseItemResponse{}, err
	}

	if !req.Amount.IsPositive() {
		return dto.ExpenseItemResponse{}, ungerr.UnprocessableEntityError(appconstant.ErrNonPositiveAmount)
	}

	request := expenseitem.AddRequest{
		ProfileID:       profileID,
		GroupExpenseID:  req.GroupExpenseID,
		ExpenseItemData: mapper.NewExpenseItemRequestToData(req),
	}

	expenseItem, err := ges.expenseItemClient.Add(ctx, request)
	if err != nil {
		return dto.ExpenseItemResponse{}, err
	}

	profileIDs := []uuid.UUID{profileID}
	profileIDs = append(profileIDs, expenseItem.ProfileIDs()...)
	namesByProfileID, err := ges.profileService.GetNames(ctx, profileIDs)
	if err != nil {
		return dto.ExpenseItemResponse{}, err
	}

	return mapper.ExpenseItemToResponse(expenseItem, profileID, namesByProfileID), nil
}

func (ges *expenseItemServiceImpl) GetDetails(ctx context.Context, groupExpenseID, expenseItemID uuid.UUID) (dto.ExpenseItemResponse, error) {
	profileID, err := util.GetProfileID(ctx)
	if err != nil {
		return dto.ExpenseItemResponse{}, err
	}

	request := expenseitem.GetDetailsRequest{
		ID:             expenseItemID,
		GroupExpenseID: groupExpenseID,
	}

	expenseItem, err := ges.expenseItemClient.GetDetails(ctx, request)
	if err != nil {
		return dto.ExpenseItemResponse{}, err
	}

	profileIDs := []uuid.UUID{profileID}
	profileIDs = append(profileIDs, expenseItem.ProfileIDs()...)
	namesByProfileID, err := ges.profileService.GetNames(ctx, profileIDs)
	if err != nil {
		return dto.ExpenseItemResponse{}, err
	}

	return mapper.ExpenseItemToResponse(expenseItem, profileID, namesByProfileID), nil
}

func (ges *expenseItemServiceImpl) Update(ctx context.Context, req dto.UpdateExpenseItemRequest) (dto.ExpenseItemResponse, error) {
	profileID, err := util.GetProfileID(ctx)
	if err != nil {
		return dto.ExpenseItemResponse{}, err
	}

	if !req.Amount.IsPositive() {
		return dto.ExpenseItemResponse{}, ungerr.UnprocessableEntityError(appconstant.ErrNonPositiveAmount)
	}

	request := expenseitem.UpdateRequest{
		ProfileID:       profileID,
		ID:              req.ID,
		GroupExpenseID:  req.GroupExpenseID,
		ExpenseItemData: mapper.UpdateExpenseItemRequestToData(req),
	}

	expenseItem, err := ges.expenseItemClient.Update(ctx, request)
	if err != nil {
		return dto.ExpenseItemResponse{}, err
	}

	profileIDs := []uuid.UUID{profileID}
	profileIDs = append(profileIDs, expenseItem.ProfileIDs()...)
	namesByProfileID, err := ges.profileService.GetNames(ctx, profileIDs)
	if err != nil {
		return dto.ExpenseItemResponse{}, err
	}

	return mapper.ExpenseItemToResponse(expenseItem, profileID, namesByProfileID), nil
}

func (ges *expenseItemServiceImpl) Remove(ctx context.Context, groupExpenseID, expenseItemID uuid.UUID) error {
	profileID, err := util.GetProfileID(ctx)
	if err != nil {
		return err
	}

	request := expenseitem.RemoveRequest{
		ProfileID:      profileID,
		ID:             expenseItemID,
		GroupExpenseID: groupExpenseID,
	}

	return ges.expenseItemClient.Remove(ctx, request)
}

package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/domain/groupexpense"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
	"github.com/itsLeonB/orcashtrator/internal/util"
	"github.com/itsLeonB/ungerr"
	"github.com/shopspring/decimal"
)

type groupExpenseServiceImpl struct {
	friendshipService  FriendshipService
	debtService        DebtService
	profileService     ProfileService
	groupExpenseClient groupexpense.GroupExpenseClient
}

func NewGroupExpenseService(
	friendshipService FriendshipService,
	debtService DebtService,
	profileService ProfileService,
	groupExpenseClient groupexpense.GroupExpenseClient,
) GroupExpenseService {
	return &groupExpenseServiceImpl{
		friendshipService,
		debtService,
		profileService,
		groupExpenseClient,
	}
}

func (ges *groupExpenseServiceImpl) CreateDraft(ctx context.Context, request dto.NewGroupExpenseRequest) (dto.GroupExpenseResponse, error) {
	userProfileID, err := util.GetProfileID(ctx)
	if err != nil {
		return dto.GroupExpenseResponse{}, err
	}

	if err := ges.validateRequest(ctx, request); err != nil {
		return dto.GroupExpenseResponse{}, err
	}

	// Default PayerProfileID to the user's profile ID if not provided
	// This is useful when the user is creating a group expense for themselves.
	if request.PayerProfileID == uuid.Nil {
		request.PayerProfileID = userProfileID
	} else {
		// Check if the payer is a friend of the user
		isFriend, _, err := ges.friendshipService.IsFriends(ctx, userProfileID, request.PayerProfileID)
		if err != nil {
			return dto.GroupExpenseResponse{}, err
		}
		if !isFriend {
			return dto.GroupExpenseResponse{}, ungerr.UnprocessableEntityError(appconstant.ErrNotFriends)
		}
	}

	groupExpense := mapper.GroupExpenseRequestToEntity(request)
	groupExpense.CreatorProfileID = userProfileID

	insertedGroupExpense, err := ges.groupExpenseClient.CreateDraft(ctx, groupExpense)
	if err != nil {
		return dto.GroupExpenseResponse{}, err
	}

	namesByProfileIDs, err := ges.profileService.GetNames(ctx, insertedGroupExpense.ProfileIDs())
	if err != nil {
		return dto.GroupExpenseResponse{}, err
	}

	return mapper.GroupExpenseToResponse(insertedGroupExpense, userProfileID, namesByProfileIDs), nil
}

func (ges *groupExpenseServiceImpl) GetAllCreated(ctx context.Context) ([]dto.GroupExpenseResponse, error) {
	userProfileID, err := util.GetProfileID(ctx)
	if err != nil {
		return nil, err
	}

	groupExpenses, err := ges.groupExpenseClient.GetAllCreated(ctx, userProfileID)
	if err != nil {
		return nil, err
	}

	profileIDs := make([]uuid.UUID, 0)
	for _, groupExpense := range groupExpenses {
		profileIDs = append(profileIDs, groupExpense.ProfileIDs()...)
	}

	namesByProfileIDs := make(map[uuid.UUID]string, len(profileIDs))
	if len(profileIDs) > 0 {
		namesByProfileIDs, err = ges.profileService.GetNames(ctx, profileIDs)
		if err != nil {
			return nil, err
		}
	}

	mapFunc := func(groupExpense groupexpense.GroupExpense) dto.GroupExpenseResponse {
		return mapper.GroupExpenseToResponse(groupExpense, userProfileID, namesByProfileIDs)
	}

	return ezutil.MapSlice(groupExpenses, mapFunc), nil
}

func (ges *groupExpenseServiceImpl) GetDetails(ctx context.Context, id uuid.UUID) (dto.GroupExpenseResponse, error) {
	userProfileID, err := util.GetProfileID(ctx)
	if err != nil {
		return dto.GroupExpenseResponse{}, err
	}

	groupExpense, err := ges.groupExpenseClient.GetDetails(ctx, id)
	if err != nil {
		return dto.GroupExpenseResponse{}, err
	}

	namesByProfileIDs, err := ges.profileService.GetNames(ctx, groupExpense.ProfileIDs())
	if err != nil {
		return dto.GroupExpenseResponse{}, err
	}

	return mapper.GroupExpenseToResponse(groupExpense, userProfileID, namesByProfileIDs), nil
}

func (ges *groupExpenseServiceImpl) ConfirmDraft(ctx context.Context, id uuid.UUID) (dto.GroupExpenseResponse, error) {
	userProfileID, err := util.GetProfileID(ctx)
	if err != nil {
		return dto.GroupExpenseResponse{}, err
	}

	request := groupexpense.ConfirmDraftRequest{
		ID:        id,
		ProfileID: userProfileID,
	}

	groupExpense, err := ges.groupExpenseClient.ConfirmDraft(ctx, request)
	if err != nil {
		return dto.GroupExpenseResponse{}, err
	}

	if err = ges.debtService.ProcessConfirmedGroupExpense(ctx, groupExpense); err != nil {
		return dto.GroupExpenseResponse{}, err
	}

	namesByProfileIDs, err := ges.profileService.GetNames(ctx, groupExpense.ProfileIDs())
	if err != nil {
		return dto.GroupExpenseResponse{}, err
	}

	return mapper.GroupExpenseToResponse(groupExpense, userProfileID, namesByProfileIDs), nil
}

func (ges *groupExpenseServiceImpl) validateRequest(ctx context.Context, request dto.NewGroupExpenseRequest) error {
	if request.TotalAmount.IsZero() {
		return ungerr.UnprocessableEntityError(appconstant.ErrAmountZero)
	}

	calculatedFeeTotal := decimal.Zero
	calculatedSubtotal := decimal.Zero
	for _, item := range request.Items {
		calculatedSubtotal = calculatedSubtotal.Add(item.Amount.Mul(decimal.NewFromInt(int64(item.Quantity))))
	}
	for _, fee := range request.OtherFees {
		calculatedFeeTotal = calculatedFeeTotal.Add(fee.Amount)
	}
	if calculatedFeeTotal.Add(calculatedSubtotal).Cmp(request.TotalAmount) != 0 {
		return ungerr.UnprocessableEntityError(appconstant.ErrAmountMismatched)
	}
	if calculatedSubtotal.Cmp(request.Subtotal) != 0 {
		return ungerr.UnprocessableEntityError(appconstant.ErrAmountMismatched)
	}

	return nil
}

package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/domain/otherfee"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
	"github.com/itsLeonB/orcashtrator/internal/util"
	"github.com/itsLeonB/ungerr"
)

type otherFeeServiceImpl struct {
	profileService ProfileService
	otherFeeClient otherfee.OtherFeeClient
}

func NewOtherFeeService(
	profileService ProfileService,
	otherFeeClient otherfee.OtherFeeClient,
) OtherFeeService {
	return &otherFeeServiceImpl{
		profileService,
		otherFeeClient,
	}
}

func (ofs *otherFeeServiceImpl) Add(ctx context.Context, req dto.NewOtherFeeRequest) (dto.OtherFeeResponse, error) {
	profileID, err := util.GetProfileID(ctx)
	if err != nil {
		return dto.OtherFeeResponse{}, err
	}

	if !req.Amount.IsPositive() {
		return dto.OtherFeeResponse{}, ungerr.UnprocessableEntityError(appconstant.ErrNonPositiveAmount)
	}

	request := otherfee.AddRequest{
		ProfileID:      profileID,
		GroupExpenseID: req.GroupExpenseID,
		OtherFeeData: otherfee.OtherFeeData{
			Name:              req.Name,
			Amount:            req.Amount,
			CalculationMethod: req.CalculationMethod,
		},
	}

	otherFee, err := ofs.otherFeeClient.Add(ctx, request)
	if err != nil {
		return dto.OtherFeeResponse{}, err
	}

	profileIDs := []uuid.UUID{profileID}
	profileIDs = append(profileIDs, otherFee.ProfileIDs()...)
	namesByProfileID, err := ofs.profileService.GetNames(ctx, profileIDs)
	if err != nil {
		return dto.OtherFeeResponse{}, err
	}

	return mapper.OtherFeeToResponse(otherFee, profileID, namesByProfileID), nil
}

func (ofs *otherFeeServiceImpl) Update(ctx context.Context, req dto.UpdateOtherFeeRequest) (dto.OtherFeeResponse, error) {
	profileID, err := util.GetProfileID(ctx)
	if err != nil {
		return dto.OtherFeeResponse{}, err
	}

	if !req.Amount.IsPositive() {
		return dto.OtherFeeResponse{}, ungerr.UnprocessableEntityError(appconstant.ErrNonPositiveAmount)
	}

	request := otherfee.UpdateRequest{
		ProfileID:      profileID,
		ID:             req.ID,
		GroupExpenseID: req.GroupExpenseID,
		OtherFeeData: otherfee.OtherFeeData{
			Name:              req.Name,
			Amount:            req.Amount,
			CalculationMethod: req.CalculationMethod,
		},
	}

	otherFee, err := ofs.otherFeeClient.Update(ctx, request)
	if err != nil {
		return dto.OtherFeeResponse{}, err
	}

	profileIDs := []uuid.UUID{profileID}
	profileIDs = append(profileIDs, otherFee.ProfileIDs()...)
	namesByProfileID, err := ofs.profileService.GetNames(ctx, profileIDs)
	if err != nil {
		return dto.OtherFeeResponse{}, err
	}

	return mapper.OtherFeeToResponse(otherFee, profileID, namesByProfileID), nil
}

func (ofs *otherFeeServiceImpl) Remove(ctx context.Context, groupExpenseID, otherFeeID uuid.UUID) error {
	profileID, err := util.GetProfileID(ctx)
	if err != nil {
		return err
	}

	request := otherfee.RemoveRequest{
		ProfileID:      profileID,
		ID:             otherFeeID,
		GroupExpenseID: groupExpenseID,
	}

	return ofs.otherFeeClient.Remove(ctx, request)
}

func (ofs *otherFeeServiceImpl) GetCalculationMethods(ctx context.Context) ([]dto.FeeCalculationMethodInfo, error) {
	methods, err := ofs.otherFeeClient.GetCalculationMethods(ctx)
	if err != nil {
		return nil, err
	}

	return ezutil.MapSlice(methods, mapper.FeeCalculationMethodInfoToResponse), nil
}

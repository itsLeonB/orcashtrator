package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/domain/expensebill"
	"github.com/itsLeonB/orcashtrator/internal/domain/expensebill/uploadbill"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
	"github.com/itsLeonB/ungerr"
)

type expenseBillServiceImpl struct {
	logger            ezutil.Logger
	friendshipService FriendshipService
	profileService    ProfileService
	expenseBillClient expensebill.ExpenseBillClient
	uploadBillClient  uploadbill.UploadBillClient
}

func NewExpenseBillService(
	logger ezutil.Logger,
	friendshipService FriendshipService,
	profileService ProfileService,
	expenseBillClient expensebill.ExpenseBillClient,
	uploadBillClient uploadbill.UploadBillClient,
) ExpenseBillService {
	return &expenseBillServiceImpl{
		logger,
		friendshipService,
		profileService,
		expenseBillClient,
		uploadBillClient,
	}
}

func (ebs *expenseBillServiceImpl) Save(ctx context.Context, req *dto.NewExpenseBillRequest) (dto.ExpenseBillResponse, error) {
	var err error
	var namesByProfileIDs map[uuid.UUID]string

	// Default PayerProfileID to the user's profile ID if not provided
	// This is useful when the user is creating a group expense for themselves.
	if req.PayerProfileID != req.CreatorProfileID {
		// Check if the payer is a friend of the user
		if isFriend, _, err := ebs.friendshipService.IsFriends(ctx, req.CreatorProfileID, req.PayerProfileID); err != nil {
			return dto.ExpenseBillResponse{}, err
		} else if !isFriend {
			return dto.ExpenseBillResponse{}, ungerr.UnprocessableEntityError(appconstant.ErrNotFriends)
		}
		namesByProfileIDs, err = ebs.profileService.GetNames(ctx, []uuid.UUID{req.CreatorProfileID, req.PayerProfileID})
		if err != nil {
			return dto.ExpenseBillResponse{}, err
		}
	} else {
		req.PayerProfileID = req.CreatorProfileID
		creatorProfile, err := ebs.profileService.GetByID(ctx, req.CreatorProfileID)
		if err != nil {
			return dto.ExpenseBillResponse{}, err
		}
		namesByProfileIDs = map[uuid.UUID]string{
			req.CreatorProfileID: creatorProfile.Name,
		}
	}

	uploadReq := uploadbill.UploadBillRequest{
		FileStream:  req.ImageReader,
		ContentType: req.ContentType,
		Filename:    req.Filename,
		FileSize:    req.FileSize,
	}

	objectKey, err := ebs.uploadBillClient.UploadStream(ctx, &uploadReq)
	if err != nil {
		return dto.ExpenseBillResponse{}, err
	}

	request := expensebill.ExpenseBill{
		CreatorProfileID: req.CreatorProfileID,
		PayerProfileID:   req.PayerProfileID,
		ObjectKey:        objectKey,
	}

	savedBill, err := ebs.expenseBillClient.Save(ctx, request)
	if err != nil {
		if err = ebs.uploadBillClient.Delete(ctx, objectKey); err != nil {
			ebs.logger.Errorf("error rolling back bill upload: %v", err)
		}
		return dto.ExpenseBillResponse{}, err
	}

	return mapper.ExpenseBillToResponse(savedBill, "", req.CreatorProfileID, namesByProfileIDs), nil
}

func (ebs *expenseBillServiceImpl) GetAllCreated(ctx context.Context, creatorProfileID uuid.UUID) ([]dto.ExpenseBillResponse, error) {
	bills, err := ebs.expenseBillClient.GetAllCreated(ctx, creatorProfileID)
	if err != nil {
		return nil, err
	}

	profileIDs := mapper.UniqueBillProfileIDs(bills)
	namesByProfileIDs, err := ebs.profileService.GetNames(ctx, profileIDs)
	if err != nil {
		return nil, err
	}

	return ezutil.MapSlice(bills, mapper.ExpenseBillSimpleMapper("", creatorProfileID, namesByProfileIDs)), nil
}

func (ebs *expenseBillServiceImpl) Get(ctx context.Context, profileID, id uuid.UUID) (dto.ExpenseBillResponse, error) {
	bill, err := ebs.expenseBillClient.Get(ctx, id)
	if err != nil {
		return dto.ExpenseBillResponse{}, err
	}

	imageURL, err := ebs.uploadBillClient.GetURL(ctx, bill.ObjectKey)
	if err != nil {
		return dto.ExpenseBillResponse{}, err
	}

	namesByProfileIDs, err := ebs.profileService.GetNames(ctx, []uuid.UUID{bill.CreatorProfileID, bill.PayerProfileID})
	if err != nil {
		return dto.ExpenseBillResponse{}, err
	}

	return mapper.ExpenseBillToResponse(bill, imageURL, profileID, namesByProfileIDs), nil
}

func (ebs *expenseBillServiceImpl) Delete(ctx context.Context, profileID, id uuid.UUID) error {
	bill, err := ebs.expenseBillClient.Get(ctx, id)
	if err != nil {
		return err
	}

	if err = ebs.expenseBillClient.Delete(ctx, profileID, id); err != nil {
		return err
	}

	if err = ebs.uploadBillClient.Delete(ctx, bill.ObjectKey); err != nil {
		ebs.logger.Errorf("error deleting bill image from storage: %v", err)
	}

	return nil
}

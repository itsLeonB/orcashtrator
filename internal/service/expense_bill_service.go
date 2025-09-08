package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/domain/expensebill"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/util"
	"github.com/itsLeonB/ungerr"
)

type expenseBillServiceImpl struct {
	logger            ezutil.Logger
	friendshipService FriendshipService
	expenseBillClient expensebill.ExpenseBillClient
}

func NewExpenseBillService(
	logger ezutil.Logger,
	friendshipService FriendshipService,
	expenseBillClient expensebill.ExpenseBillClient,
) ExpenseBillService {
	return &expenseBillServiceImpl{
		logger,
		friendshipService,
		expenseBillClient,
	}
}

func (ebs *expenseBillServiceImpl) Upload(ctx context.Context, req dto.NewExpenseBillRequest) (dto.UploadBillResponse, error) {
	defer func() {
		if err := req.ImageReader.Close(); err != nil {
			ebs.logger.Errorf("error closing ImageReader: %v\n", err)
		}
	}()

	userProfileID, err := util.GetProfileID(ctx)
	if err != nil {
		return dto.UploadBillResponse{}, err
	}

	// Default PayerProfileID to the user's profile ID if not provided
	// This is useful when the user is creating a group expense for themselves.
	if req.PayerProfileID == uuid.Nil {
		req.PayerProfileID = userProfileID
	} else {
		// Check if the payer is a friend of the user
		if isFriend, _, err := ebs.friendshipService.IsFriends(ctx, userProfileID, req.PayerProfileID); err != nil {
			return dto.UploadBillResponse{}, err
		} else if !isFriend {
			return dto.UploadBillResponse{}, ungerr.UnprocessableEntityError(appconstant.ErrNotFriends)
		}
	}

	request := expensebill.UploadStreamRequest{
		CreatorProfileID: userProfileID,
		PayerProfileID:   req.PayerProfileID,
		FileStream:       req.ImageReader,
	}

	id, err := ebs.expenseBillClient.UploadStream(ctx, request)
	if err != nil {
		return dto.UploadBillResponse{}, err
	}

	return dto.UploadBillResponse{ID: id}, nil
}

package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/domain/expensebill"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/ungerr"
)

type expenseBillServiceImpl struct {
	friendshipService FriendshipService
	expenseBillClient expensebill.ExpenseBillClient
}

func NewExpenseBillService(
	friendshipService FriendshipService,
	expenseBillClient expensebill.ExpenseBillClient,
) ExpenseBillService {
	return &expenseBillServiceImpl{
		friendshipService,
		expenseBillClient,
	}
}

func (ebs *expenseBillServiceImpl) Upload(ctx context.Context, req dto.NewExpenseBillRequest) (dto.UploadBillResponse, error) {
	// Default PayerProfileID to the user's profile ID if not provided
	// This is useful when the user is creating a group expense for themselves.
	if req.PayerProfileID == uuid.Nil {
		req.PayerProfileID = req.CreatorProfileID
	} else if req.PayerProfileID != req.CreatorProfileID {
		// Check if the payer is a friend of the user
		if isFriend, _, err := ebs.friendshipService.IsFriends(ctx, req.CreatorProfileID, req.PayerProfileID); err != nil {
			return dto.UploadBillResponse{}, err
		} else if !isFriend {
			return dto.UploadBillResponse{}, ungerr.UnprocessableEntityError(appconstant.ErrNotFriends)
		}
	}

	request := expensebill.UploadStreamRequest{
		CreatorProfileID: req.CreatorProfileID,
		PayerProfileID:   req.PayerProfileID,
		FileStream:       req.ImageReader,
		ContentType:      req.ContentType,
		Filename:         req.Filename,
		FileSize:         req.FileSize,
	}

	id, err := ebs.expenseBillClient.UploadStream(ctx, request)
	if err != nil {
		return dto.UploadBillResponse{}, err
	}

	return dto.UploadBillResponse{ID: id}, nil
}

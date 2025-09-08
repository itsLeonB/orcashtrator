package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/domain/friendship"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
)

type friendshipServiceImpl struct {
	friendshipClient friendship.FriendshipClient
}

func NewFriendshipService(friendshipClient friendship.FriendshipClient) FriendshipService {
	return &friendshipServiceImpl{friendshipClient}
}

func (fs *friendshipServiceImpl) CreateAnonymous(ctx context.Context, req dto.NewAnonymousFriendshipRequest) (dto.FriendshipResponse, error) {
	request := friendship.CreateAnonymousRequest{
		ProfileID: req.ProfileID,
		Name:      req.Name,
	}

	response, err := fs.friendshipClient.CreateAnonymous(ctx, request)
	if err != nil {
		return dto.FriendshipResponse{}, err
	}

	return mapper.FriendshipToResponse(response), nil
}

func (fs *friendshipServiceImpl) GetAll(ctx context.Context, profileID uuid.UUID) ([]dto.FriendshipResponse, error) {
	response, err := fs.friendshipClient.GetAll(ctx, profileID)
	if err != nil {
		return nil, err
	}

	return ezutil.MapSlice(response, mapper.FriendshipToResponse), nil
}

func (fs *friendshipServiceImpl) IsFriends(ctx context.Context, profileID1, profileID2 uuid.UUID) (bool, bool, error) {
	request := friendship.IsFriendsRequest{
		ProfileID1: profileID1,
		ProfileID2: profileID2,
	}

	response, err := fs.friendshipClient.IsFriends(ctx, request)
	if err != nil {
		return false, false, err
	}

	return response.IsFriends, response.IsAnonymous, nil
}

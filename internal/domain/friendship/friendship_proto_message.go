package friendship

import "github.com/google/uuid"

type CreateAnonymousRequest struct {
	ProfileID uuid.UUID `validate:"required"`
	Name      string    `validate:"required,min=3"`
}

type GetDetailsRequest struct {
	ProfileID    uuid.UUID `validate:"required"`
	FriendshipID uuid.UUID `validate:"required"`
}

type IsFriendsRequest struct {
	ProfileID1 uuid.UUID `validate:"required"`
	ProfileID2 uuid.UUID `validate:"required"`
}

type IsFriendsResponse struct {
	IsFriends   bool
	IsAnonymous bool
}

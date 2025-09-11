package profile

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/cocoon-protos/gen/go/profile/v1"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/rotisserie/eris"
)

func fromProfileProto(p *profile.Profile) (Profile, error) {
	if p == nil {
		return Profile{}, eris.New("profile is nil")
	}

	userID, err := ezutil.Parse[uuid.UUID](p.GetUserId())
	if err != nil {
		return Profile{}, err
	}

	amp := p.GetAuditMetadata()
	if amp == nil {
		return Profile{}, eris.New("audit metadata is nil")
	}

	id, err := ezutil.Parse[uuid.UUID](amp.GetId())
	if err != nil {
		return Profile{}, err
	}

	return Profile{
		ID:          id,
		UserID:      userID,
		Name:        p.GetName(),
		IsAnonymous: p.IsAnonymous,
		CreatedAt:   ezutil.FromProtoTime(amp.GetCreatedAt()),
		UpdatedAt:   ezutil.FromProtoTime(amp.GetUpdatedAt()),
		DeletedAt:   ezutil.FromProtoTime(amp.GetDeletedAt()),
	}, nil
}

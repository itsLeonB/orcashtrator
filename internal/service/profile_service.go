package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/domain/profile"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
)

type profileServiceGrpc struct {
	profileClient profile.ProfileClient
}

func NewProfileService(
	profileClient profile.ProfileClient,
) ProfileService {
	return &profileServiceGrpc{
		profileClient,
	}
}

func (ps *profileServiceGrpc) GetByID(ctx context.Context, id uuid.UUID) (dto.ProfileResponse, error) {
	profile, err := ps.profileClient.Get(ctx, id)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	return mapper.ProfileToResponse(profile), nil
}

func (ps *profileServiceGrpc) GetNames(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]string, error) {
	profiles, err := ps.profileClient.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	namesByProfileID := make(map[uuid.UUID]string, len(profiles))
	for _, profile := range profiles {
		namesByProfileID[profile.ID] = profile.Name
	}

	return namesByProfileID, nil
}

func (ps *profileServiceGrpc) Update(ctx context.Context, id uuid.UUID, name string) (dto.ProfileResponse, error) {
	request := profile.UpdateRequest{
		ID:   id,
		Name: name,
	}

	updatedProfile, err := ps.profileClient.Update(ctx, request)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	return mapper.ProfileToResponse(updatedProfile), nil
}

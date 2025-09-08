package mapper

import (
	"github.com/itsLeonB/orcashtrator/internal/domain/profile"
	"github.com/itsLeonB/orcashtrator/internal/dto"
)

func ProfileToResponse(response profile.Profile) dto.ProfileResponse {
	return dto.ProfileResponse{
		ID:          response.ID,
		UserID:      response.UserID,
		Name:        response.Name,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
		DeletedAt:   response.DeletedAt,
		IsAnonymous: response.IsAnonymous,
	}
}

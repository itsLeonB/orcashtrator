package mapper

import (
	"github.com/itsLeonB/orcashtrator/internal/domain/debt"
	"github.com/itsLeonB/orcashtrator/internal/dto"
)

func TransferMethodToResponse(tm debt.TransferMethod) dto.TransferMethodResponse {
	return dto.TransferMethodResponse{
		ID:        tm.ID,
		Name:      tm.Name,
		Display:   tm.Name,
		CreatedAt: tm.CreatedAt,
		UpdatedAt: tm.UpdatedAt,
		DeletedAt: tm.DeletedAt,
	}
}

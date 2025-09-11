package service

import (
	"context"

	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/domain/debt"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
)

type transferMethodServiceImpl struct {
	transferMethodClient debt.TransferMethodClient
}

func NewTransferMethodService(transferMethodClient debt.TransferMethodClient) TransferMethodService {
	return &transferMethodServiceImpl{transferMethodClient}
}

func (tms *transferMethodServiceImpl) GetAll(ctx context.Context) ([]dto.TransferMethodResponse, error) {
	transferMethods, err := tms.transferMethodClient.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return ezutil.MapSlice(transferMethods, mapper.TransferMethodToResponse), nil
}

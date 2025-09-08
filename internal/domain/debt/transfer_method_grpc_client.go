package debt

import (
	"context"

	"github.com/itsLeonB/drex-protos/gen/go/transaction/v1"
	"github.com/itsLeonB/ezutil/v2"
	"google.golang.org/grpc"
)

type TransferMethodClient interface {
	GetAll(ctx context.Context) ([]TransferMethod, error)
}

type transferMethodClient struct {
	client transaction.TransferMethodServiceClient
}

func NewTransferMethodClient(conn *grpc.ClientConn) TransferMethodClient {
	return &transferMethodClient{transaction.NewTransferMethodServiceClient(conn)}
}

func (tc *transferMethodClient) GetAll(ctx context.Context) ([]TransferMethod, error) {
	response, err := tc.client.GetAll(ctx, nil)
	if err != nil {
		return nil, err
	}

	return ezutil.MapSliceWithError(response.GetTransferMethods(), fromTransferMethodProto)
}

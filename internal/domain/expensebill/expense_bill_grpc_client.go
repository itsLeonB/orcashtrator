package expensebill

import (
	"context"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/itsLeonB/billsplittr-protos/gen/go/expensebill/v1"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/rotisserie/eris"
	"google.golang.org/grpc"
)

type ExpenseBillClient interface {
	UploadStream(ctx context.Context, req UploadStreamRequest) (uuid.UUID, error)
}

type expenseBillClient struct {
	validate *validator.Validate
	client   expensebill.ExpenseBillServiceClient
}

func NewExpenseBillClient(
	validate *validator.Validate,
	conn *grpc.ClientConn,
) ExpenseBillClient {
	if validate == nil {
		panic("validator is nil")
	}

	return &expenseBillClient{
		validate,
		expensebill.NewExpenseBillServiceClient(conn),
	}
}

func (ebc *expenseBillClient) UploadStream(ctx context.Context, req UploadStreamRequest) (uuid.UUID, error) {
	if err := ebc.validate.Struct(req); err != nil {
		return uuid.Nil, err
	}

	stream, err := ebc.client.UploadStream(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	metadata := &expensebill.UploadStreamRequest{
		Data: &expensebill.UploadStreamRequest_BillMetadata{
			BillMetadata: &expensebill.BillMetadata{
				CreatorProfileId: req.CreatorProfileID.String(),
				PayerProfileId:   req.PayerProfileID.String(),
				ContentType:      req.ContentType,
				Filename:         req.Filename,
				FileSize:         req.FileSize,
			},
		},
	}

	if err = stream.Send(metadata); err != nil {
		return uuid.Nil, err
	}

	buffer := make([]byte, appconstant.ChunkSize)

	for {
		n, err := req.FileStream.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return uuid.Nil, eris.Wrap(err, "failed to read from stream")
		}

		chunk := &expensebill.UploadStreamRequest{
			Data: &expensebill.UploadStreamRequest_Chunk{
				Chunk: buffer[:n],
			},
		}

		if err = stream.Send(chunk); err != nil {
			return uuid.Nil, err
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		return uuid.Nil, err
	}

	return ezutil.Parse[uuid.UUID](response.GetId())
}

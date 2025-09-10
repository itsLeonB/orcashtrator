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
		return uuid.Nil, eris.Wrap(err, appconstant.ErrStructValidation)
	}

	stream, err := ebc.client.UploadStream(ctx)
	if err != nil {
		return uuid.Nil, eris.Wrap(err, "error opening grpc client stream")
	}

	if err = stream.Send(toBillMetadataProto(req)); err != nil {
		return uuid.Nil, eris.Wrap(err, "error sending metadata to grpc stream")
	}

	if err = ebc.sendDataChunks(stream, req.FileStream, req.FileSize); err != nil {
		return uuid.Nil, err
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		return uuid.Nil, eris.Wrap(err, "error closing grpc stream")
	}

	return ezutil.Parse[uuid.UUID](response.GetId())
}

func (ebc *expenseBillClient) sendDataChunks(
	stream grpc.ClientStreamingClient[expensebill.UploadStreamRequest, expensebill.UploadStreamResponse],
	fileStream io.ReadCloser,
	fileSize int64,
) error {
	buffer := make([]byte, appconstant.ChunkSize)

	var sent int64

	for {
		n, err := fileStream.Read(buffer)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return eris.Wrap(err, "failed to read from stream")
		}

		if n > 0 {
			chunk := &expensebill.UploadStreamRequest{
				Data: &expensebill.UploadStreamRequest_Chunk{
					Chunk: buffer[:n],
				},
			}
			if err = stream.Send(chunk); err != nil {
				return eris.Wrap(err, "send chunk")
			}
			sent += int64(n)
			if sent > fileSize {
				return eris.Errorf("read more bytes than declared: sent=%d declared=%d", sent, fileSize)
			}
		}
	}
}

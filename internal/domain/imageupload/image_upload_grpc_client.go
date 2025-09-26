package imageupload

import (
	"context"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/stortr-protos/gen/go/genericupload/v1"
	"github.com/itsLeonB/stortr-protos/gen/go/imageupload/v1"
	"github.com/rotisserie/eris"
	"google.golang.org/grpc"
)

type ImageUploadClient interface {
	UploadStream(ctx context.Context, req *ImageUploadRequest) (string, error)
	GetURL(ctx context.Context, fileID FileIdentifier) (string, error)
	Delete(ctx context.Context, fileID FileIdentifier) error
}

type uploadBillClient struct {
	validate *validator.Validate
	client   imageupload.ImageUploadServiceClient
}

func NewImageUploadClient(validate *validator.Validate, conn *grpc.ClientConn) ImageUploadClient {
	if validate == nil {
		panic("validate is nil")
	}

	return &uploadBillClient{
		validate,
		imageupload.NewImageUploadServiceClient(conn),
	}
}

func (ubc *uploadBillClient) UploadStream(ctx context.Context, req *ImageUploadRequest) (string, error) {
	if err := ubc.validate.Struct(req); err != nil {
		return "", eris.Wrap(err, appconstant.ErrStructValidation)
	}

	stream, err := ubc.client.UploadStream(ctx)
	if err != nil {
		return "", eris.Wrap(err, "error opening grpc client stream")
	}

	metadata, err := toMetadataProto(req)
	if err != nil {
		return "", err
	}

	if err = stream.Send(metadata); err != nil {
		return "", eris.Wrap(err, "error sending data to grpc stream")
	}

	if err = ubc.sendDataChunks(stream, req.FileStream, req.FileSize); err != nil {
		return "", err
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		return "", eris.Wrap(err, "error closing grpc stream")
	}

	return response.GetUri(), nil
}

func (ubc *uploadBillClient) GetURL(ctx context.Context, fileID FileIdentifier) (string, error) {
	request := genericupload.GetUrlRequest{
		FileIdentifier: toFileIdentifierProto(fileID),
	}

	response, err := ubc.client.GetUrl(ctx, &request)
	if err != nil {
		return "", err
	}

	return response.GetUrl(), nil
}

func (ubc *uploadBillClient) Delete(ctx context.Context, fileID FileIdentifier) error {
	request := genericupload.DeleteRequest{
		FileIdentifier: toFileIdentifierProto(fileID),
	}

	_, err := ubc.client.Delete(ctx, &request)

	return err
}

func (ubc *uploadBillClient) sendDataChunks(
	stream grpc.ClientStreamingClient[genericupload.UploadStreamRequest, genericupload.UploadStreamResponse],
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
			chunk := &genericupload.UploadStreamRequest{
				Data: &genericupload.UploadStreamRequest_Chunk{
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

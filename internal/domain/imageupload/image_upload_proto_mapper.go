package imageupload

import (
	"github.com/itsLeonB/stortr-protos/gen/go/genericupload/v1"
	"github.com/rotisserie/eris"
)

func toMetadataProto(req *ImageUploadRequest) (*genericupload.UploadStreamRequest, error) {
	if req == nil {
		return nil, eris.New("upload bill request is nil")
	}

	return &genericupload.UploadStreamRequest{
		Data: &genericupload.UploadStreamRequest_Metadata{
			Metadata: &genericupload.Metadata{
				ContentType: req.ContentType,
				FileSize:    req.FileSize,
				FileIdentifier: &genericupload.FileIdentifier{
					BucketName: req.BucketName,
					ObjectKey:  req.ObjectKey,
				},
			},
		},
	}, nil
}

func toFileIdentifierProto(req FileIdentifier) *genericupload.FileIdentifier {
	return &genericupload.FileIdentifier{
		BucketName: req.BucketName,
		ObjectKey:  req.ObjectKey,
	}
}

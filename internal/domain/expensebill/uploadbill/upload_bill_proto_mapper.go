package uploadbill

import (
	"github.com/itsLeonB/stortr-protos/gen/go/uploadbill/v1"
	"github.com/rotisserie/eris"
)

func toBillMetadataProto(req *UploadBillRequest) (*uploadbill.UploadStreamRequest, error) {
	if req == nil {
		return nil, eris.New("upload bill request is nil")
	}

	return &uploadbill.UploadStreamRequest{
		Data: &uploadbill.UploadStreamRequest_BillMetadata{
			BillMetadata: &uploadbill.BillMetadata{
				ContentType: req.ContentType,
				Filename:    req.Filename,
				FileSize:    req.FileSize,
			},
		},
	}, nil
}

package uploadbill

import (
	"github.com/itsLeonB/stortr-protos/gen/go/expensebill/v1"
	"github.com/rotisserie/eris"
)

func toBillMetadataProto(req *UploadBillRequest) (*expensebill.UploadStreamRequest, error) {
	if req == nil {
		return nil, eris.New("upload bill request is nil")
	}

	return &expensebill.UploadStreamRequest{
		Data: &expensebill.UploadStreamRequest_BillMetadata{
			BillMetadata: &expensebill.BillMetadata{
				ContentType: req.ContentType,
				Filename:    req.Filename,
				FileSize:    req.FileSize,
			},
		},
	}, nil
}

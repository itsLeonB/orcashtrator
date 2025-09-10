package expensebill

import "github.com/itsLeonB/billsplittr-protos/gen/go/expensebill/v1"

func toBillMetadataProto(req UploadStreamRequest) *expensebill.UploadStreamRequest {
	return &expensebill.UploadStreamRequest{
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
}

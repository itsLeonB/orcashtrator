package uploadbill

import "io"

type UploadBillRequest struct {
	FileStream  io.ReadCloser `validate:"required"`
	ContentType string        `validate:"required,oneof=image/jpeg image/png image/jpg image/webp"`
	Filename    string        `validate:"required,min=3"`
	FileSize    int64         `validate:"required,min=1"`
}

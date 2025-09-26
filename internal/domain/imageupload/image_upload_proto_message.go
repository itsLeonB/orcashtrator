package imageupload

import "io"

type ImageUploadRequest struct {
	FileStream  io.ReadCloser `validate:"required"`
	ContentType string        `validate:"required,oneof=image/jpeg image/png image/jpg image/webp"`
	FileSize    int64         `validate:"required,min=1"`
	FileIdentifier
}

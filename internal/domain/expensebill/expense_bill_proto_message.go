package expensebill

import (
	"io"

	"github.com/google/uuid"
)

type UploadStreamRequest struct {
	CreatorProfileID uuid.UUID     `validate:"required"`
	PayerProfileID   uuid.UUID     `validate:"required"`
	FileStream       io.ReadCloser `validate:"required"`
	ContentType      string        `validate:"required,oneof=image/jpeg image/jpg image/png image/webp"`
	Filename         string        `validate:"required,min=3"`
	FileSize         int64         `validate:"required,min=1"`
}

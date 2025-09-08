package dto

import (
	"io"

	"github.com/google/uuid"
)

type NewExpenseBillRequest struct {
	PayerProfileID uuid.UUID
	ImageReader    io.ReadCloser
	ContentType    string
	Filename       string
	FileSize       int64
}

type UploadBillResponse struct {
	ID uuid.UUID `json:"id"`
}

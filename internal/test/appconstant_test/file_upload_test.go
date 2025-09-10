package appconstant_test

import (
	"testing"

	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/stretchr/testify/assert"
)

func TestFileUploadConstants(t *testing.T) {
	assert.Equal(t, 64*1024, appconstant.ChunkSize)
	assert.Equal(t, 10*1024*1024, appconstant.MaxFileSize)
}

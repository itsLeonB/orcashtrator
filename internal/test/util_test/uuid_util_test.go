package util_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	id := uuid.New()
	result := util.ToString(id)
	assert.Equal(t, id.String(), result)
}

package mapper_test

import (
	"testing"

	"github.com/itsLeonB/orcashtrator/internal/domain/otherfee"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
	"github.com/stretchr/testify/assert"
)

func TestFeeCalculationMethodInfoToResponse(t *testing.T) {
	cmi := otherfee.CalculationMethodInfo{
		Name:        "EQUAL_SPLIT",
		Display:     "Equal Split",
		Description: "Split the fee equally among all participants",
	}

	result := mapper.FeeCalculationMethodInfoToResponse(cmi)

	assert.Equal(t, cmi.Name, result.Name)
	assert.Equal(t, cmi.Display, result.Display)
	assert.Equal(t, cmi.Description, result.Description)
}

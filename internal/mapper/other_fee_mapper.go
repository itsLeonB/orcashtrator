package mapper

import (
	"github.com/itsLeonB/orcashtrator/internal/domain/otherfee"
	"github.com/itsLeonB/orcashtrator/internal/dto"
)

func FeeCalculationMethodInfoToResponse(cmi otherfee.CalculationMethodInfo) dto.FeeCalculationMethodInfo {
	return dto.FeeCalculationMethodInfo{
		Name:        cmi.Name,
		Display:     cmi.Display,
		Description: cmi.Description,
	}
}

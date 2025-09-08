package otherfee

import (
	"github.com/itsLeonB/billsplittr-protos/gen/go/otherfee/v1"
)

type CalculationMethodInfo struct {
	Name        string
	Display     string
	Description string
}

func fromCalculationMethodInfoProto(cmi *otherfee.FeeCalculationMethodInfo) CalculationMethodInfo {
	return CalculationMethodInfo{
		Name:        cmi.GetMethod().String(),
		Display:     cmi.GetDisplay(),
		Description: cmi.GetDescription(),
	}
}

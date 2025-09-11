package otherfee

import (
	"github.com/itsLeonB/billsplittr-protos/gen/go/otherfee/v1"
	"github.com/rotisserie/eris"
)

type CalculationMethodInfo struct {
	Name        string
	Display     string
	Description string
}

func fromCalculationMethodInfoProto(cmi *otherfee.FeeCalculationMethodInfo) (CalculationMethodInfo, error) {
	if cmi == nil {
		return CalculationMethodInfo{}, eris.New("fee calculation method info proto struct is nil")
	}

	return CalculationMethodInfo{
		Name:        cmi.GetMethod().String(),
		Display:     cmi.GetDisplay(),
		Description: cmi.GetDescription(),
	}, nil
}

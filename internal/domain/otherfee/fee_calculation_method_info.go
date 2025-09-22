package otherfee

import (
	"github.com/itsLeonB/billsplittr-protos/gen/go/otherfee/v1"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/rotisserie/eris"
)

type CalculationMethodInfo struct {
	Name        appconstant.FeeCalculationMethod
	Display     string
	Description string
}

func fromCalculationMethodInfoProto(cmi *otherfee.FeeCalculationMethodInfo) (CalculationMethodInfo, error) {
	if cmi == nil {
		return CalculationMethodInfo{}, eris.New("fee calculation method info proto struct is nil")
	}

	name, err := fromCalculationMethodEnum(cmi.GetMethod())
	if err != nil {
		return CalculationMethodInfo{}, err
	}

	return CalculationMethodInfo{
		Name:        name,
		Display:     cmi.GetDisplay(),
		Description: cmi.GetDescription(),
	}, nil
}

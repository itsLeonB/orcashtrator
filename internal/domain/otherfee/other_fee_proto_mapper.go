package otherfee

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/billsplittr-protos/gen/go/otherfee/v1"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/domain"
	"github.com/rotisserie/eris"
	"golang.org/x/text/currency"
)

func ToOtherFeeProto(of OtherFeeData) (*otherfee.OtherFee, error) {
	method, err := toCalculationMethodEnum(of.CalculationMethod)
	if err != nil {
		return nil, err
	}
	return &otherfee.OtherFee{
		Name:              of.Name,
		Amount:            ezutil.DecimalToMoney(of.Amount, currency.IDR.String()),
		CalculationMethod: method,
	}, nil
}

func FromOtherFeeResponseProto(ofr *otherfee.OtherFeeResponse) (OtherFee, error) {
	if ofr == nil {
		return OtherFee{}, eris.New("other fee response is nil")
	}

	groupExpenseID, err := ezutil.Parse[uuid.UUID](ofr.GetGroupExpenseId())
	if err != nil {
		return OtherFee{}, err
	}

	fee := ofr.GetOtherFee()
	if fee == nil {
		return OtherFee{}, eris.New("expense item data is nil")
	}

	participants, err := ezutil.MapSliceWithError(ofr.GetParticipants(), fromFeeParticipantProto)
	if err != nil {
		return OtherFee{}, err
	}

	metadata, err := domain.FromAuditMetadataProto(ofr.GetAuditMetadata())
	if err != nil {
		return OtherFee{}, err
	}

	calcMethod, err := fromCalculationMethodEnum(fee.GetCalculationMethod())
	if err != nil {
		return OtherFee{}, err
	}

	return OtherFee{
		GroupExpenseID: groupExpenseID,
		OtherFeeData: OtherFeeData{
			Name:              fee.GetName(),
			Amount:            ezutil.MoneyToDecimal(fee.GetAmount()),
			CalculationMethod: calcMethod,
		},
		Participants:  participants,
		AuditMetadata: metadata,
	}, nil
}

func fromFeeParticipantProto(fpr *otherfee.FeeParticipantResponse) (FeeParticipant, error) {
	if fpr == nil {
		return FeeParticipant{}, eris.New("item participant is nil")
	}

	profileID, err := ezutil.Parse[uuid.UUID](fpr.GetProfileId())
	if err != nil {
		return FeeParticipant{}, err
	}

	return FeeParticipant{
		ProfileID:   profileID,
		ShareAmount: ezutil.MoneyToDecimal(fpr.GetShareAmount()),
	}, nil
}

func fromCalculationMethodEnum(cm otherfee.FeeCalculationMethod) (appconstant.FeeCalculationMethod, error) {
	switch cm {
	case otherfee.FeeCalculationMethod_FEE_CALCULATION_METHOD_EQUAL_SPLIT:
		return appconstant.EqualSplitFee, nil
	case otherfee.FeeCalculationMethod_FEE_CALCULATION_METHOD_ITEMIZED_SPLIT:
		return appconstant.ItemizedSplitFee, nil
	default:
		return "", eris.Errorf("unknown fee calculation method enum: %s", cm.String())
	}
}

func toCalculationMethodEnum(cm appconstant.FeeCalculationMethod) (otherfee.FeeCalculationMethod, error) {
	switch cm {
	case appconstant.EqualSplitFee:
		return otherfee.FeeCalculationMethod_FEE_CALCULATION_METHOD_EQUAL_SPLIT, nil
	case appconstant.ItemizedSplitFee:
		return otherfee.FeeCalculationMethod_FEE_CALCULATION_METHOD_ITEMIZED_SPLIT, nil
	default:
		return otherfee.FeeCalculationMethod_FEE_CALCULATION_METHOD_UNSPECIFIED, eris.Errorf("unknown fee calculation method constant: %s", cm)
	}
}

package otherfee

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/domain"
	"github.com/shopspring/decimal"
)

type OtherFee struct {
	GroupExpenseID uuid.UUID
	OtherFeeData
	Participants []FeeParticipant
	domain.AuditMetadata
}

func (of OtherFee) ProfileIDs() []uuid.UUID {
	return ezutil.MapSlice(of.Participants, func(fp FeeParticipant) uuid.UUID { return fp.ProfileID })
}

type OtherFeeData struct {
	Name              string                           `validate:"required,min=3"`
	Amount            decimal.Decimal                  `validate:"required"`
	CalculationMethod appconstant.FeeCalculationMethod `validate:"required"`
}

type FeeParticipant struct {
	ProfileID   uuid.UUID
	ShareAmount decimal.Decimal
}

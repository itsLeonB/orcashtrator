package debt

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/drex-protos/gen/go/debt/v1"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/rotisserie/eris"
	"golang.org/x/text/currency"
)

func toTransactionActionEnum(ta appconstant.Action) (debt.TransactionAction, error) {
	switch ta {
	case appconstant.BorrowAction:
		return debt.TransactionAction_TRANSACTION_ACTION_BORROW, nil
	case appconstant.LendAction:
		return debt.TransactionAction_TRANSACTION_ACTION_LEND, nil
	case appconstant.ReceiveAction:
		return debt.TransactionAction_TRANSACTION_ACTION_RECEIVE, nil
	case appconstant.ReturnAction:
		return debt.TransactionAction_TRANSACTION_ACTION_RETURN, nil
	default:
		return debt.TransactionAction_TRANSACTION_ACTION_UNSPECIFIED, eris.Errorf("undefined TransactionAction constant: %s", ta)
	}
}

func fromTransactionActionEnum(ta debt.TransactionAction) (appconstant.Action, error) {
	switch ta {
	case debt.TransactionAction_TRANSACTION_ACTION_BORROW:
		return appconstant.BorrowAction, nil
	case debt.TransactionAction_TRANSACTION_ACTION_LEND:
		return appconstant.LendAction, nil
	case debt.TransactionAction_TRANSACTION_ACTION_RECEIVE:
		return appconstant.ReceiveAction, nil
	case debt.TransactionAction_TRANSACTION_ACTION_RETURN:
		return appconstant.ReturnAction, nil
	default:
		return "", eris.Errorf("undefined TransactionAction enum: %s", ta)
	}
}

func fromTransactionTypeEnum(ta debt.TransactionType) (appconstant.DebtTransactionType, error) {
	switch ta {
	case debt.TransactionType_TRANSACTION_TYPE_LEND:
		return appconstant.Lend, nil
	case debt.TransactionType_TRANSACTION_TYPE_REPAY:
		return appconstant.Repay, nil
	default:
		return "", eris.Errorf("undefined TransactionType enum: %s", ta)
	}
}

func fromTransactionProto(trx *debt.TransactionResponse) (Transaction, error) {
	if trx == nil {
		return Transaction{}, eris.New("transaction from response is nil")
	}

	id, err := ezutil.Parse[uuid.UUID](trx.GetId())
	if err != nil {
		return Transaction{}, err
	}

	profileID, err := ezutil.Parse[uuid.UUID](trx.GetProfileId())
	if err != nil {
		return Transaction{}, err
	}

	trxType, err := fromTransactionTypeEnum(trx.GetType())
	if err != nil {
		return Transaction{}, err
	}

	trxAction, err := fromTransactionActionEnum(trx.GetAction())
	if err != nil {
		return Transaction{}, err
	}

	return Transaction{
		ID:             id,
		ProfileID:      profileID,
		Type:           trxType,
		Action:         trxAction,
		Amount:         ezutil.MoneyToDecimal(trx.GetAmount()),
		TransferMethod: trx.GetTransferMethod(),
		Description:    trx.GetDescription(),
		CreatedAt:      ezutil.FromProtoTime(trx.GetCreatedAt()),
		UpdatedAt:      ezutil.FromProtoTime(trx.GetUpdatedAt()),
		DeletedAt:      ezutil.FromProtoTime(trx.GetDeletedAt()),
	}, nil
}

func toExpenseParticipantProto(ep ExpenseParticipantData) *debt.ExpenseParticipantData {
	return &debt.ExpenseParticipantData{
		ProfileId:   ep.ProfileID.String(),
		ShareAmount: ezutil.DecimalToMoney(ep.ShareAmount, currency.IDR.String()),
	}
}

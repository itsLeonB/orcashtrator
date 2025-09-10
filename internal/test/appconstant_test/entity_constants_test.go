package appconstant_test

import (
	"testing"

	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/stretchr/testify/assert"
)

func TestDebtTransactionType(t *testing.T) {
	assert.Equal(t, appconstant.DebtTransactionType("LEND"), appconstant.Lend)
	assert.Equal(t, appconstant.DebtTransactionType("REPAY"), appconstant.Repay)
}

func TestFriendshipType(t *testing.T) {
	assert.Equal(t, appconstant.FriendshipType("REAL"), appconstant.Real)
	assert.Equal(t, appconstant.FriendshipType("ANON"), appconstant.Anonymous)
}

func TestFeeCalculationMethod(t *testing.T) {
	assert.Equal(t, appconstant.FeeCalculationMethod("EQUAL_SPLIT"), appconstant.EqualSplitFee)
	assert.Equal(t, appconstant.FeeCalculationMethod("ITEMIZED_SPLIT"), appconstant.ItemizedSplitFee)
}

func TestGroupExpenseTransferMethod(t *testing.T) {
	assert.Equal(t, "GROUP_EXPENSE", appconstant.GroupExpenseTransferMethod)
}

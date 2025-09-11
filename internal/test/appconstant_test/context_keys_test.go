package appconstant_test

import (
	"testing"

	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/stretchr/testify/assert"
)

func TestContextKeys(t *testing.T) {
	assert.Equal(t, "userID", appconstant.ContextUserID.String())
	assert.Equal(t, "profileID", appconstant.ContextProfileID.String())
	assert.Equal(t, "friendshipID", appconstant.ContextFriendshipID.String())
	assert.Equal(t, "groupExpenseID", appconstant.ContextGroupExpenseID.String())
	assert.Equal(t, "expenseItemID", appconstant.ContextExpenseItemID.String())
	assert.Equal(t, "otherFeeID", appconstant.ContextOtherFeeID.String())
}

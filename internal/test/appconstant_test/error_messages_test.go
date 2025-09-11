package appconstant_test

import (
	"testing"

	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/stretchr/testify/assert"
)

func TestErrorMessages(t *testing.T) {
	assert.Equal(t, "error retrieving data", appconstant.ErrDataSelect)
	assert.Equal(t, "error inserting new data", appconstant.ErrDataInsert)
	assert.Equal(t, "error updating data", appconstant.ErrDataUpdate)
	assert.Equal(t, "error deleting data", appconstant.ErrDataDelete)
	assert.Equal(t, "user is not found", appconstant.ErrAuthUserNotFound)
	assert.Equal(t, "user with email %s is already registered", appconstant.ErrAuthDuplicateUser)
	assert.Equal(t, "unknown credentials, please check your email/password", appconstant.ErrAuthUnknownCredentials)
	assert.Equal(t, "user with ID: %s is not found", appconstant.ErrUserNotFound)
	assert.Equal(t, "user with ID: %s is deleted", appconstant.ErrUserDeleted)
	assert.Equal(t, "friendship not found", appconstant.ErrFriendshipNotFound)
	assert.Equal(t, "friendship is deleted", appconstant.ErrFriendshipDeleted)
	assert.Equal(t, "amount mismatch, please check the total amount and the items/fees provided", appconstant.ErrAmountMismatched)
	assert.Equal(t, "amount must be greater than zero", appconstant.ErrAmountZero)
	assert.Equal(t, "you are not friends with this user, please add them as a friend first", appconstant.ErrNotFriends)
	assert.Equal(t, "error processing file upload", appconstant.ErrProcessFile)
	assert.Equal(t, "amount must be positive (>0)", appconstant.ErrNonPositiveAmount)
	assert.Equal(t, "service client communication failure", appconstant.ErrServiceClient)
	assert.Equal(t, "error validating struct input", appconstant.ErrStructValidation)
}

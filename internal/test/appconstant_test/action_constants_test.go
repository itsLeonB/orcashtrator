package appconstant_test

import (
	"testing"

	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/stretchr/testify/assert"
)

func TestAction(t *testing.T) {
	assert.Equal(t, appconstant.Action("LEND"), appconstant.LendAction)
	assert.Equal(t, appconstant.Action("BORROW"), appconstant.BorrowAction)
	assert.Equal(t, appconstant.Action("RECEIVE"), appconstant.ReceiveAction)
	assert.Equal(t, appconstant.Action("RETURN"), appconstant.ReturnAction)
}

package appconstant_test

import (
	"testing"

	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/stretchr/testify/assert"
)

func TestResponseMessages(t *testing.T) {
	assert.Equal(t, "success retrieving data", appconstant.MsgGetData)
	assert.Equal(t, "success inserting new data", appconstant.MsgInsertData)
	assert.Equal(t, "success updating data", appconstant.MsgUpdateData)
	assert.Equal(t, "success registering, please login", appconstant.MsgRegisterSuccess)
	assert.Equal(t, "success login", appconstant.MsgLoginSuccess)
	assert.Equal(t, "bill uploaded", appconstant.MsgBillUploaded)
}

package mapper

import (
	"github.com/itsLeonB/orcashtrator/internal/domain/auth"
	"github.com/itsLeonB/orcashtrator/internal/dto"
)

func LoginToResponse(login auth.LoginResponse) dto.LoginResponse {
	return dto.LoginResponse{
		Type:  login.Type,
		Token: login.Token,
	}
}

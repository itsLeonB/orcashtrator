package auth

import "github.com/itsLeonB/cocoon-protos/gen/go/auth/v1"

func fromLoginResponseProto(resp *auth.LoginResponse) LoginResponse {
	return LoginResponse{
		Type:  resp.GetType(),
		Token: resp.GetToken(),
	}
}

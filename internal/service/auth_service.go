package service

import (
	"context"

	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/domain/auth"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/rotisserie/eris"
)

type authServiceGrpc struct {
	authClient      auth.AuthClient
	verificationURL string
}

func NewAuthService(
	authClient auth.AuthClient,
	verificationURL string,
) AuthService {
	return &authServiceGrpc{
		authClient,
		verificationURL,
	}
}

func (as *authServiceGrpc) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	request := auth.RegisterRequest{
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
		VerificationURL:      as.verificationURL,
	}

	isVerified, err := as.authClient.Register(ctx, request)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	msg := "check your email to confirm your registration"
	if isVerified {
		msg = "success registering, please login"
	}

	return dto.RegisterResponse{
		Message: msg,
	}, nil
}

func (as *authServiceGrpc) InternalLogin(ctx context.Context, req dto.InternalLoginRequest) (dto.LoginResponse, error) {
	request := auth.InternalLoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	response, err := as.authClient.InternalLogin(ctx, request)
	if err != nil {
		return dto.LoginResponse{}, eris.Wrap(err, appconstant.ErrServiceClient)
	}

	return dto.LoginResponse{
		Type:  response.Type,
		Token: response.Token,
	}, nil
}

func (as *authServiceGrpc) VerifyToken(ctx context.Context, token string) (bool, map[string]any, error) {
	return as.authClient.VerifyToken(ctx, token)
}

func (as *authServiceGrpc) GetOAuth2URL(ctx context.Context, provider string) (string, error) {
	return as.authClient.GetOAuth2URL(ctx, provider)
}

func (as *authServiceGrpc) OAuth2Login(ctx context.Context, provider, code, state string) (dto.LoginResponse, error) {
	request := auth.OAuthLoginRequest{
		Provider: provider,
		Code:     code,
		State:    state,
	}

	response, err := as.authClient.OAuth2Login(ctx, request)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		Type:  response.Type,
		Token: response.Token,
	}, nil
}

func (as *authServiceGrpc) VerifyRegistration(ctx context.Context, token string) (dto.LoginResponse, error) {
	response, err := as.authClient.VerifyRegistration(ctx, token)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		Type:  response.Type,
		Token: response.Token,
	}, nil
}

package service

import (
	"context"

	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/domain/auth"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/rotisserie/eris"
)

type authServiceGrpc struct {
	authClient auth.AuthClient
}

func NewAuthService(
	authClient auth.AuthClient,
) AuthService {
	return &authServiceGrpc{
		authClient,
	}
}

func (as *authServiceGrpc) Register(ctx context.Context, req dto.RegisterRequest) error {
	request := auth.RegisterRequest{
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	if err := as.authClient.Register(ctx, request); err != nil {
		return err
	}

	return nil
}

func (as *authServiceGrpc) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	request := auth.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	response, err := as.authClient.Login(ctx, request)
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

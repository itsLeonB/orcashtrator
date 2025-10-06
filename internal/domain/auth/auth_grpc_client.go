package auth

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/itsLeonB/cocoon-protos/gen/go/auth/v1"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/ungerr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthClient interface {
	Register(ctx context.Context, req RegisterRequest) error
	InternalLogin(ctx context.Context, req InternalLoginRequest) (LoginResponse, error)
	OAuth2Login(ctx context.Context, req OAuthLoginRequest) (LoginResponse, error)
	VerifyToken(ctx context.Context, token string) (bool, map[string]any, error)
	GetOAuth2URL(ctx context.Context, provider string) (string, error)
}

type authClient struct {
	validate *validator.Validate
	client   auth.AuthServiceClient
}

func NewAuthClient(validate *validator.Validate, conn *grpc.ClientConn) AuthClient {
	if validate == nil {
		panic("validator is nil")
	}

	return &authClient{
		validate: validate,
		client:   auth.NewAuthServiceClient(conn),
	}
}

func (ac *authClient) Register(ctx context.Context, req RegisterRequest) error {
	if err := ac.validate.Struct(req); err != nil {
		return err
	}

	request := &auth.RegisterRequest{
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	if _, err := ac.client.Register(ctx, request); err != nil {
		return err
	}

	return nil
}

func (ac *authClient) InternalLogin(ctx context.Context, req InternalLoginRequest) (LoginResponse, error) {
	if err := ac.validate.Struct(req); err != nil {
		return LoginResponse{}, err
	}

	request := auth.LoginRequest{
		InternalRequest: &auth.InternalLoginRequest{
			Email:    req.Email,
			Password: req.Password,
		},
	}

	response, err := ac.client.Login(ctx, &request)
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		Type:  response.GetType(),
		Token: response.GetToken(),
	}, nil
}

func (ac *authClient) OAuth2Login(ctx context.Context, req OAuthLoginRequest) (LoginResponse, error) {
	if err := ac.validate.Struct(req); err != nil {
		return LoginResponse{}, err
	}

	request := auth.LoginRequest{
		Oauth2Request: &auth.OAuth2LoginRequest{
			Provider: req.Provider,
			Code:     req.Code,
			State:    req.State,
		},
	}

	response, err := ac.client.Login(ctx, &request)
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		Type:  response.GetType(),
		Token: response.GetToken(),
	}, nil
}

func (ac *authClient) VerifyToken(ctx context.Context, token string) (bool, map[string]any, error) {
	if token == "" {
		return false, nil, ungerr.BadRequestError("token is empty")
	}

	request := &auth.VerifyTokenRequest{Token: token}

	data, err := ac.client.VerifyToken(ctx, request)
	if err != nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.Unauthenticated {
			return false, nil, ungerr.UnauthorizedError("unauthorized")
		}
		return false, nil, err
	}

	return true, map[string]any{
		appconstant.ContextProfileID.String(): data.GetProfileId(),
	}, nil
}

func (ac *authClient) GetOAuth2URL(ctx context.Context, provider string) (string, error) {
	if provider == "" {
		return "", ungerr.BadRequestError("provider is empty")
	}

	request := auth.GetOAuth2UrlRequest{Provider: provider}
	response, err := ac.client.GetOAuth2Url(ctx, &request)
	if err != nil {
		return "", err
	}

	return response.GetUrl(), nil
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/itsLeonB/ginkgo"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/service"
	"github.com/itsLeonB/ungerr"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(
	authService service.AuthService,
) *AuthHandler {
	return &AuthHandler{
		authService,
	}
}

func (ah *AuthHandler) HandleRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request, err := ginkgo.BindRequest[dto.RegisterRequest](ctx, binding.JSON)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		response, err := ah.authService.Register(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(
			http.StatusCreated,
			ginkgo.NewResponse("success registering").WithData(response),
		)
	}
}

func (ah *AuthHandler) HandleInternalLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request, err := ginkgo.BindRequest[dto.InternalLoginRequest](ctx, binding.JSON)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		response, err := ah.authService.InternalLogin(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(
			http.StatusOK,
			ginkgo.NewResponse(appconstant.MsgLoginSuccess).WithData(response),
		)
	}
}

func (ah *AuthHandler) HandleOAuth2Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		provider, err := ah.getProvider(ctx)
		if err != nil {
			_ = ctx.Error(ungerr.BadRequestError("missing oauth provider"))
			return
		}

		url, err := ah.authService.GetOAuth2URL(ctx, provider)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func (ah *AuthHandler) HandleOAuth2Callback() gin.HandlerFunc {
	return ginkgo.WrapHandler(func(ctx *gin.Context) (int, string, any, error) {
		provider, err := ah.getProvider(ctx)
		if err != nil {
			return 0, "", nil, err
		}
		code := ctx.Query("code")
		state := ctx.Query("state")

		response, err := ah.authService.OAuth2Login(ctx, provider, code, state)
		if err != nil {
			return 0, "", nil, err
		}

		return http.StatusOK, "success logging in", response, nil
	})
}

func (ah *AuthHandler) HandleVerifyRegistration() gin.HandlerFunc {
	return ginkgo.WrapHandler(func(ctx *gin.Context) (int, string, any, error) {
		token := ctx.Query("token")
		if token == "" {
			return 0, "", nil, ungerr.BadRequestError("missing token")
		}

		response, err := ah.authService.VerifyRegistration(ctx, token)
		if err != nil {
			return 0, "", nil, err
		}

		return http.StatusOK, "registration verified", response, nil
	})
}

func (ah *AuthHandler) HandleSendPasswordReset() gin.HandlerFunc {
	return ginkgo.WrapHandler(func(ctx *gin.Context) (int, string, any, error) {
		request, err := ginkgo.BindRequest[dto.SendPasswordResetRequest](ctx, binding.JSON)
		if err != nil {
			return 0, "", nil, err
		}

		if err = ah.authService.SendPasswordReset(ctx, request.Email); err != nil {
			return 0, "", nil, err
		}

		return http.StatusCreated, "reset password link sent to email", nil, nil
	})
}

func (ah *AuthHandler) HandleResetPassword() gin.HandlerFunc {
	return ginkgo.WrapHandler(func(ctx *gin.Context) (int, string, any, error) {
		request, err := ginkgo.BindRequest[dto.ResetPasswordRequest](ctx, binding.JSON)
		if err != nil {
			return 0, "", nil, err
		}

		response, err := ah.authService.ResetPassword(ctx, request.Token, request.Password)
		if err != nil {
			return 0, "", nil, err
		}

		return http.StatusOK, "password successfully reset", response, nil
	})
}

func (ah *AuthHandler) getProvider(ctx *gin.Context) (string, error) {
	provider := ctx.Param(appconstant.ContextProvider.String())
	if provider == "" {
		return "", ungerr.BadRequestError("missing oauth provider")
	}
	return provider, nil
}

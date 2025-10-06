package handler

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/itsLeonB/ginkgo"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/service"
	"github.com/itsLeonB/ungerr"
	"github.com/rotisserie/eris"
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

		err = ah.authService.Register(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(
			http.StatusCreated,
			ginkgo.NewResponse(appconstant.MsgRegisterSuccess),
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

		state, err := generateStateToken()
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.SetCookie("oauth_state", state, 300, "/", "", true, true)

		url, err := ah.authService.GetOAuth2URL(ctx, provider, state)
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

		cookieState, err := ctx.Cookie("oauth_state")
		if err != nil || cookieState != state {
			return 0, "", nil, eris.Wrap(err, "invalid state parameter")
		}

		ctx.SetCookie("oauth_state", "", -1, "/", "", true, true)

		response, err := ah.authService.OAuth2Login(ctx, provider, code, state)
		if err != nil {
			return 0, "", nil, err
		}

		return http.StatusOK, "success logging in", response, nil
	})
}

func (ah *AuthHandler) getProvider(ctx *gin.Context) (string, error) {
	provider := ctx.Param(appconstant.ContextProvider.String())
	if provider == "" {
		return "", ungerr.BadRequestError("missing oauth provider")
	}
	return provider, nil
}

func generateStateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", eris.Wrap(err, "error generating random string")
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

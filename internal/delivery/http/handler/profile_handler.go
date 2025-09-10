package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/itsLeonB/ginkgo"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/service"
)

type ProfileHandler struct {
	profileService service.ProfileService
}

func NewProfileHandler(
	profileService service.ProfileService,
) *ProfileHandler {
	return &ProfileHandler{
		profileService,
	}
}

func (ph *ProfileHandler) HandleProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		profileID, err := ginkgo.GetAndParseFromContext[uuid.UUID](ctx, appconstant.ContextProfileID.String())
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		response, err := ph.profileService.GetByID(ctx, profileID)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(
			http.StatusOK,
			ginkgo.NewResponse(appconstant.MsgGetData).WithData(response),
		)
	}
}

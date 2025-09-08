package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/itsLeonB/ginkgo"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/service"
)

type OtherFeeHandler struct {
	otherFeeSvc service.OtherFeeService
}

func NewOtherFeeHandler(
	otherFeeSvc service.OtherFeeService,
) *OtherFeeHandler {
	return &OtherFeeHandler{
		otherFeeSvc,
	}
}

func (geh *OtherFeeHandler) HandleAdd() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		groupExpenseID, err := ginkgo.GetRequiredPathParam[uuid.UUID](ctx, appconstant.ContextGroupExpenseID)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		request, err := ginkgo.BindRequest[dto.NewOtherFeeRequest](ctx, binding.JSON)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		request.GroupExpenseID = groupExpenseID

		response, err := geh.otherFeeSvc.Add(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(
			http.StatusCreated,
			ginkgo.NewResponse(appconstant.MsgInsertData).WithData(response),
		)
	}
}

func (geh *OtherFeeHandler) HandleUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		groupExpenseID, err := ginkgo.GetRequiredPathParam[uuid.UUID](ctx, appconstant.ContextGroupExpenseID)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		otherFeeID, err := ginkgo.GetRequiredPathParam[uuid.UUID](ctx, appconstant.ContextOtherFeeID)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		request, err := ginkgo.BindRequest[dto.UpdateOtherFeeRequest](ctx, binding.JSON)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		request.GroupExpenseID = groupExpenseID
		request.ID = otherFeeID

		response, err := geh.otherFeeSvc.Update(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(
			http.StatusOK,
			ginkgo.NewResponse(appconstant.MsgUpdateData).WithData(response),
		)
	}
}

func (geh *OtherFeeHandler) HandleRemove() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		groupExpenseID, err := ginkgo.GetRequiredPathParam[uuid.UUID](ctx, appconstant.ContextGroupExpenseID)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		feeID, err := ginkgo.GetRequiredPathParam[uuid.UUID](ctx, appconstant.ContextOtherFeeID)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		if err = geh.otherFeeSvc.Remove(ctx, groupExpenseID, feeID); err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}

func (geh *OtherFeeHandler) HandleGetFeeCalculationMethods() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := geh.otherFeeSvc.GetCalculationMethods(ctx)
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

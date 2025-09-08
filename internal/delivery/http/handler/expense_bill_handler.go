package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/ginkgo"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/service"
	"github.com/rotisserie/eris"
)

type ExpenseBillHandler struct {
	expenseBillService service.ExpenseBillService
}

func NewExpenseBillHandler(
	expenseBillService service.ExpenseBillService,
) *ExpenseBillHandler {
	return &ExpenseBillHandler{
		expenseBillService,
	}
}

func (geh *ExpenseBillHandler) HandleUploadBill() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		payerProfileID := uuid.Nil
		if payerProfileIDStr := ctx.PostForm("payerProfileId"); payerProfileIDStr != "" {
			payerProfileID, err = ezutil.Parse[uuid.UUID](payerProfileIDStr)
			if err != nil {
				_ = ctx.Error(err)
				return
			}
		}

		fileHeader, err := ctx.FormFile("bill")
		if err != nil {
			_ = ctx.Error(eris.Wrap(err, appconstant.ErrProcessFile))
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			_ = ctx.Error(eris.Wrap(err, appconstant.ErrProcessFile))
			return
		}

		request := dto.NewExpenseBillRequest{
			PayerProfileID: payerProfileID,
			ImageReader:    file,
			ContentType:    fileHeader.Header.Get("Content-Type"),
			Filename:       fileHeader.Filename,
			FileSize:       fileHeader.Size,
		}

		response, err := geh.expenseBillService.Upload(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(
			http.StatusCreated,
			ginkgo.NewResponse(appconstant.MsgBillUploaded).WithData(response),
		)
	}
}

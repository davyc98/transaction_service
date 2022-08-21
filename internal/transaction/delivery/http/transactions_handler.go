package http

import (
	"context"
	"net/http"
	models "transaction_service/internal/domain/model"
	"transaction_service/internal/domain/request"
	"transaction_service/internal/domain/response"
	"transaction_service/internal/transaction/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	transactionUseCase *usecase.TransactionUseCase
	ctx                context.Context
}

func NewHandler(ctx context.Context, transactionUseCase *usecase.TransactionUseCase) *Handler {
	return &Handler{ctx: ctx, transactionUseCase: transactionUseCase}
}

func (h Handler) GetMerchantTransaction(ctx echo.Context) error {
	var req request.TransactionRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	if req.Limit == 0 {
		req.Limit = -1
	}
	paging, err := h.transactionUseCase.GetMerchantsTransaction(ctx, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.APIResponse{
			Message: err.Error(),
		})
	}

	transactions := paging.Records.(*[]models.TransactionsQuery)
	return ctx.JSON(http.StatusOK, response.APIResponse{
		Message:  "successfully get all merchants data",
		Data:     transactions,
		PageInfo: response.NewPageInfo().ToPageInfo(paging),
	})
}

func (h Handler) GetOutletTransaction(ctx echo.Context) error {
	var req request.TransactionRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	if req.Limit == 0 {
		req.Limit = -1
	}
	paging, err := h.transactionUseCase.GetOutletsTransaction(ctx, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.APIResponse{
			Message: err.Error(),
		})
	}

	transactions := paging.Records.(*[]models.TransactionsQuery)
	return ctx.JSON(http.StatusOK, response.APIResponse{
		Message:  "successfully get all merchants data",
		Data:     transactions,
		PageInfo: response.NewPageInfo().ToPageInfo(paging),
	})
}

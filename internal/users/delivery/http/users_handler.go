package http

import (
	"context"
	"net/http"
	"transaction_service/internal/domain/request"
	"transaction_service/internal/domain/response"
	"transaction_service/internal/users/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	userUseCase *usecase.UserUseCase
	ctx         context.Context
}

func NewHandler(ctx context.Context, userUseCase *usecase.UserUseCase) *Handler {
	return &Handler{ctx: ctx, userUseCase: userUseCase}
}

func (h Handler) Login(ctx echo.Context) error {
	var req request.LoginUserRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	tokenDetails, err := h.userUseCase.Login(h.ctx, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.APIResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.APIResponse{
		Message: "successfully login to service",
		Data:    tokenDetails,
	})
}

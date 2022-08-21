package usecase

import (
	models "transaction_service/internal/domain/model"
	"transaction_service/internal/domain/request"

	"github.com/labstack/echo"
)

type TransactionUsecase interface {
	GetMerchants(ctx echo.Context, req request.TransactionRequest) ([]models.Transactions, error)
}

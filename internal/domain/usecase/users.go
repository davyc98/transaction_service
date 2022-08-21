package usecase

import (
	models "transaction_service/internal/domain/model"
	"transaction_service/internal/domain/request"

	"github.com/labstack/echo"
)

type UsersUsecase interface {
	Login(ctx echo.Context, req request.LoginUserRequest) (models.TokenDetails, error)
}

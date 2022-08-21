package routers

import (
	"context"
	transactionHandler "transaction_service/internal/transaction/delivery/http"
	transactionRepository "transaction_service/internal/transaction/repository/postgres"
	transactionUsecase "transaction_service/internal/transaction/usecase"
	userHandler "transaction_service/internal/users/delivery/http"
	userRepository "transaction_service/internal/users/repository/postgres"
	userUsecase "transaction_service/internal/users/usecase"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func InitAuthRouters(ctx context.Context, db *gorm.DB, app *echo.Echo) {
	userRepository := userRepository.NewUserPostgresRepository(db)
	userUseCase := userUsecase.NewUserUseCase(ctx, userRepository)
	userHTTPHandler := userHandler.NewHandler(ctx, userUseCase)
	app.POST("/login", userHTTPHandler.Login)
}

func InitTransactionRouters(ctx context.Context, db *gorm.DB, r *echo.Group) {
	transactionRepository := transactionRepository.NewTransactionPostgresRepository(db)
	transactionUseCase := transactionUsecase.NewTransactionUseCase(ctx, transactionRepository)
	transactionHTTPHandler := transactionHandler.NewHandler(ctx, transactionUseCase)
	r.GET("/merchants/:selectMonth", transactionHTTPHandler.GetMerchantTransaction)
	r.GET("/outlets/:selectMonth", transactionHTTPHandler.GetOutletTransaction)
}

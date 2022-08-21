package usecase

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"
	"transaction_service/internal/commons"
	models "transaction_service/internal/domain/model"
	"transaction_service/internal/domain/repository"
	"transaction_service/internal/domain/request"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type TransactionUseCase struct {
	transactionRepository repository.TransactionRepository
	ctx                   context.Context
}

func NewTransactionUseCase(
	ctx context.Context,
	transactionRepository repository.TransactionRepository,
) *TransactionUseCase {
	return &TransactionUseCase{
		transactionRepository: transactionRepository,
		ctx:                   ctx,
	}
}

func (u *TransactionUseCase) GetMerchantsTransaction(ctx echo.Context, req request.TransactionRequest) (paging *commons.Pagination, err error) {
	userID, err := u.getUserID(ctx)
	if err != nil {
		return nil, err
	}
	selectMonth := ctx.Param("selectMonth")
	splitSelect := strings.Split(selectMonth, "-")
	fromMonth, toMonth, err := getFromDateAndToDate(splitSelect[1], splitSelect[0])
	if err != nil {
		return nil, err
	}
	paging, err = u.transactionRepository.GetAllMerchants(ctx.Request().Context(), req.Page, req.Limit, userID, fromMonth, toMonth)
	if err != nil {
		return nil, err
	}
	return paging, nil
}

func (u *TransactionUseCase) GetOutletsTransaction(ctx echo.Context, req request.TransactionRequest) (paging *commons.Pagination, err error) {
	userID, err := u.getUserID(ctx)
	if err != nil {
		return nil, err
	}
	selectMonth := ctx.Param("selectMonth")
	splitSelect := strings.Split(selectMonth, "-")
	fromMonth, toMonth, err := getFromDateAndToDate(splitSelect[1], splitSelect[0])
	if err != nil {
		return nil, err
	}
	paging, err = u.transactionRepository.GetAllOutlets(ctx.Request().Context(), req.Page, req.Limit, userID, fromMonth, toMonth)
	if err != nil {
		return nil, err
	}
	return paging, nil
}

func (u *TransactionUseCase) getUserID(ctx echo.Context) (userID int, err error) {
	token := ctx.Get("user").(*jwt.Token)
	if token == nil {
		err = errors.New("Failed to get user token")
		return userID, err
	}
	userToken := token.Claims.(*models.JWTClaims)
	return userToken.UserID, nil
}

func getFromDateAndToDate(inputMonth string, inputYear string) (fromMonth time.Time, toMonth time.Time, err error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	month, err := strconv.Atoi(inputMonth)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	year, err := strconv.Atoi(inputYear)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	fromMonth = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, location)
	toMonth = time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, location)
	return
}

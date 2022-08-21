package repository

import (
	"context"
	"time"
	"transaction_service/internal/commons"
)

type TransactionRepository interface {
	GetAllMerchants(ctx context.Context, page, limit, userID int, fromMonth, toMonth time.Time) (*commons.Pagination, error)
	GetAllOutlets(ctx context.Context, page, limit, outletID int, fromMonth, toMonth time.Time) (*commons.Pagination, error)
}

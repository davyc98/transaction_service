package postgres

import (
	"context"
	"time"
	"transaction_service/internal/commons"
	models "transaction_service/internal/domain/model"
	"transaction_service/internal/domain/repository"

	"github.com/jinzhu/gorm"
)

type transactionPostgresRepository struct {
	db *gorm.DB
}

func NewTransactionPostgresRepository(db *gorm.DB) repository.TransactionRepository {
	return &transactionPostgresRepository{db: db}
}

func (r transactionPostgresRepository) GetAllMerchants(ctx context.Context, page, limit, userID int, fromMonth, toMonth time.Time) (*commons.Pagination, error) {
	var transactions []models.TransactionsQuery
	query := r.db.Model(&models.TransactionsQuery{}).Select("A.merchant_name, bill_total").
		Joins("JOIN merchants AS A ON a.id = transactions.merchant_id").
		Joins("JOIN users AS u ON u.id = A.user_id").
		Where("user_id = ?", userID).
		Where("transactions.created_at BETWEEN ? AND ?", fromMonth, toMonth).Find(&transactions)
	paging, err := commons.NewPagination().Paging(query, page, limit, []string{"transactions.created_at"}, &transactions)
	if err != nil {
		return nil, err
	}
	return paging, nil
}

func (r transactionPostgresRepository) FindMerchantByUserID(ctx context.Context, userID int) (models.Merchants, error) {
	var merchant models.Merchants
	err := r.db.Model(&models.Merchants{}).Where("user_id = ?", userID).Find(&merchant).Error
	return merchant, err
}

func (r transactionPostgresRepository) GetAllOutlets(ctx context.Context, page, limit, userID int, fromMonth, toMonth time.Time) (*commons.Pagination, error) {
	var transactions []models.TransactionsQuery
	query := r.db.Model(&models.TransactionsQuery{}).Select("A.merchant_name,b.outlet_name, bill_total").
		Joins("JOIN merchants AS A ON a.id = transactions.merchant_id").
		Joins("JOIN outlets AS B ON b.id = transactions.outlet_id").
		Joins("JOIN users AS u ON u.id = A.user_id").
		Where("user_id = ?", userID).
		Where("transactions.created_at BETWEEN ? AND ?", fromMonth, toMonth).Find(&transactions)
	paging, err := commons.NewPagination().Paging(query, page, limit, []string{"transactions.created_at"}, &transactions)
	if err != nil {
		return nil, err
	}
	return paging, nil
}

func (r transactionPostgresRepository) FindOutletByUserID(ctx context.Context, userID int) (models.Outlets, error) {
	var outlet models.Outlets
	err := r.db.Model(&models.Outlets{}).Where("user_id = ?", userID).Find(&outlet).Error
	return outlet, err
}

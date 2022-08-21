package postgres

import (
	"context"
	models "transaction_service/internal/domain/model"
	"transaction_service/internal/domain/repository"

	"github.com/jinzhu/gorm"
)

type userPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) repository.UserRepository {
	return &userPostgresRepository{db: db}
}

func (r userPostgresRepository) Login(ctx context.Context, user models.User) error {
	err := r.db.Model(&models.User{}).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r userPostgresRepository) FindUserByEmail(ctx context.Context, userName string) (models.User, error) {
	var user models.User
	query := r.db.Model(&models.User{}).Where("user_name = ?", userName).Find(&user)
	if query.Error != nil {
		return user, query.Error
	}
	return user, nil
}

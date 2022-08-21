package repository

import (
	"context"
	models "transaction_service/internal/domain/model"
)

type UserRepository interface {
	Login(ctx context.Context, user models.User) error
	FindUserByEmail(ctx context.Context, userName string) (models.User, error)
}

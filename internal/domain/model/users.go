package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID        int        `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	UserName  string     `gorm:"column:user_name" json:"user_name"`
	Password  string     `gorm:"column:password" json:"password"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy string     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy string     `gorm:"column:updated_by" json:"updated_by"`
}

func (User) TableName() string {
	return "users"
}

type TokenDetails struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
	UserID      int    `json:"user_id"`
}

type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

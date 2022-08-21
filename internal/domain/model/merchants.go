package models

import "time"

type Merchants struct {
	ID           int        `gorm:"column:id" json:"id"`
	UserID       int        `gorm:"column:user_id" json:"user_id"`
	MerchantName string     `gorm:"column:merchant_name" json:"merchant_name"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy    string     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy    string     `gorm:"column:updated_by" json:"updated_by"`
}

func (Merchants) TableName() string {
	return "merchants"
}

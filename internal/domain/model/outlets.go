package models

import "time"

type Outlets struct {
	ID         int        `gorm:"column:id" json:"id"`
	MerchantID int        `gorm:"column:merchant_id" json:"merchant_id"`
	OutletName string     `gorm:"column:outlet_name" json:"outlet_name"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy  string     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy  string     `gorm:"column:updated_by" json:"updated_by"`
}

func (Outlets) TableName() string {
	return "outlets"
}

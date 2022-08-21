package models

import "time"

type Transactions struct {
	ID         int        `gorm:"column:id" json:"id"`
	MerchantID int        `gorm:"column:merchant_id" json:"merchant_id"`
	OutletID   int        `gorm:"column:outlet_id" json:"outlet_id"`
	BillTotal  float64    `gorm:"column:bill_total" json:"bill_total"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy  string     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy  string     `gorm:"column:updated_by" json:"updated_by"`
}

func (Transactions) TableName() string {
	return "transactions"
}

type TransactionsQuery struct {
	MerchantID   int     `gorm:"column:merchant_id" json:"merchant_id,omitempty"`
	MerchantName string  `gorm:"column:merchant_name" json:"merchant_name"`
	OutletName   string  `gorm:"column:outlet_name" json:"outlet_name,omitempty"`
	OutletID     int     `gorm:"column:outlet_id" json:"outlet_id,omitempty"`
	BillTotal    float64 `gorm:"column:bill_total" json:"bill_total"`
}

func (TransactionsQuery) TableName() string {
	return "transactions"
}

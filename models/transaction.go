package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID              uint            `gorm:"column:id"`
	TransactionId   string          `gorm:"column:transaction_id"`
	WalletId        string          `gorm:"column:wallet_id"`
	Name            string          `gorm:"column:name"`
	Amount          decimal.Decimal `gorm:"column:amount"`
	Note            string          `gorm:"column:note"`
	Type            string          `gorm:"column:type"`
	TransactionDate time.Time       `gorm:"column:transaction_date"`
	CategoryId      string          `gorm:"column:category_id"`
	UserId          string          `gorm:"column:user_id"`
	CreatedAt       time.Time       `gorm:"column:created_at"`
	UpdatedAt       time.Time       `gorm:"column:updated_at"`
}

func (Transaction) TableName() string {
	return "transactions"
}

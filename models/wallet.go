package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID           uint            `gorm:"column:id"`
	WalletId     string          `gorm:"column:wallet_id"`
	WalletName   string          `gorm:"column:wallet_name"`
	WalletDetail *string         `gorm:"column:wallet_detail"`
	Balance      decimal.Decimal `gorm:"column:balance"`
	UserId       string          `gorm:"column:user_id"`
	CreatedAt    time.Time       `gorm:"column:created_at"`
	UpdatedAt    time.Time       `gorm:"column:updated_at"`
}

func (Wallet) TableName() string {
	return "wallets"
}

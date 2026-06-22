package response

import (
	"github.com/shopspring/decimal"
)

type TransactionResponse struct {
	TransactionId   string
	WalletId        string
	Name            string
	Amount          decimal.Decimal
	Note            string
	Type            string
	TransactionDate string `json:"transaction_date"`
	Category        CategoryResponse
	UserId          string
}

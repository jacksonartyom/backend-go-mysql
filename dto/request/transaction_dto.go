package request

import "github.com/shopspring/decimal"

type TransactionDto struct {
	WalletId        string
	Name            string
	Amount          decimal.Decimal
	Note            string
	Type            string
	TransactionDate string
	CategoryId      string
	UserId          string
}

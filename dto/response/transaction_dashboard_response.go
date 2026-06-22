package response

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransactionDashboardResponse struct {
	TransactionId   string
	WalletId        string
	Name            string
	Amount          decimal.Decimal
	Note            string
	Type            string
	TransactionDate time.Time
	CategoryId      string
	CategoryName    string
	UserId          string
}

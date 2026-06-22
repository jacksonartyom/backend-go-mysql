package response

import "github.com/shopspring/decimal"

type DashboardResponse struct {
	TotalBalance       decimal.Decimal `json:"total_balance"`
	Wallets            []WalletResponse
	RecentTransactions []TransactionDashboardResponse `json:"recent_transactions"`
}

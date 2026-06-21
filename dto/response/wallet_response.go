package response

import "github.com/shopspring/decimal"

type WalletResponse struct {
	WalletId     string          `json:"_id"`
	WalletName   string          `json:"wallet_name"`
	WalletDetail *string         `json:"wallet_detail"`
	Balance      decimal.Decimal `json:"balance"`
	UserId       string          `json:"userId"`
}

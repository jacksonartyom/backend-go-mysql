package request

import "github.com/shopspring/decimal"

type WalletDto struct {
	WalletName   string
	WalletDetail string
	Balance      decimal.Decimal
}

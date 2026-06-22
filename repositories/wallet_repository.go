package repositories

import (
	"backend-go-mysql/config"
	"backend-go-mysql/models"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type WalletRepository struct{}

func NewWalletRepository() WalletRepository {
	return WalletRepository{}
}

func (r *WalletRepository) FindWalletByUserId(userId string) ([]models.Wallet, error) {
	var wallet []models.Wallet
	result := config.DB.Where("user_id = ?", userId).Find(&wallet)

	return wallet, result.Error
}

func (r *WalletRepository) CreateWallet(wallet models.Wallet) (models.Wallet, error) {
	result := config.DB.Create(&wallet)
	return wallet, result.Error
}

func (r *WalletRepository) UpdateWallet(walletId string, wallet *models.Wallet) error {
	return config.DB.Model(&models.Wallet{}).
		Where("wallet_id = ?", walletId).
		Updates(map[string]interface{}{
			"wallet_name":   wallet.WalletName,
			"wallet_detail": wallet.WalletDetail,
			"updated_at":    wallet.UpdatedAt,
		}).Error
}

func (r *WalletRepository) DeleteWalletByWalletId(walletId string) error {
	return config.DB.
		Where("wallet_id = ?", walletId).
		Delete(&models.Wallet{}).Error
}

func (r *WalletRepository) UpdateBalance(walletId string, amount decimal.Decimal) error {
	return config.DB.
		Model(&models.Wallet{}).
		Where("wallet_id = ?", walletId).
		Update("balance", gorm.Expr("balance + ?", amount)).Error
}

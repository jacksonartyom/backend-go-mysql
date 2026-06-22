package repositories

import (
	"backend-go-mysql/config"
	"backend-go-mysql/dto/response"
	"backend-go-mysql/models"
	"time"
)

type TransactionRepository struct{}

func NewTransactionRepository() TransactionRepository {
	return TransactionRepository{}
}

func (r *TransactionRepository) CreateTransaction(transactions []models.Transaction) error {
	return config.DB.Create(&transactions).Error
}

func (r *TransactionRepository) FindByWalletId(
	walletId string,
	dateFrom time.Time,
	dateTo time.Time,
) ([]response.TransactionDashboardResponse, error) {

	var results []response.TransactionDashboardResponse

	query := `
	SELECT 
		tran.transaction_id,
		tran.wallet_id,
		tran.name,
		tran.amount,
		tran.note,
		tran.type,
		tran.transaction_date,
		tran.category_id,
		cate.name AS category_name,
		tran.user_id
	FROM transactions tran
	LEFT JOIN categories cate 
		ON tran.category_id = cate.category_id
	WHERE tran.wallet_id = ?
		AND tran.transaction_date >= ?
		AND tran.transaction_date < ?
	ORDER BY tran.transaction_date DESC
	`

	err := config.DB.Raw(query, walletId, dateFrom, dateTo).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *TransactionRepository) FindTop5ByUserId(userId string) ([]response.TransactionDashboardResponse, error) {

	var results []response.TransactionDashboardResponse

	err := config.DB.
		Table("transactions tran").
		Select(`
			tran.transaction_id,
			tran.wallet_id,
			tran.name,
			tran.amount,
			tran.note,
			tran.type,
			tran.transaction_date,
			tran.category_id,
			cate.name AS category_name,
			tran.user_id
		`).
		Joins("LEFT JOIN categories cate ON tran.category_id = cate.category_id").
		Where("tran.user_id = ?", userId).
		Order("tran.transaction_date DESC").
		Limit(5).
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}

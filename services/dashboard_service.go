package services

import (
	"backend-go-mysql/dto/response"
	"backend-go-mysql/repositories"
)

type DashboardService struct {
	TransactionRepo repositories.TransactionRepository
	WalletRepo      repositories.WalletRepository
}

func NewDashboardService(transactionRepo repositories.TransactionRepository, walletRepo repositories.WalletRepository) DashboardService {
	return DashboardService{TransactionRepo: transactionRepo,
		WalletRepo: walletRepo}
}

func (s *DashboardService) GetCategoryByUserId(userId string) (response.DashboardResponse, error) {
	totalBalance, err := s.WalletRepo.SumBalanceByUserId(userId)
	wallet, err := s.WalletRepo.FindWalletByUserId(userId)
	transactions, err := s.TransactionRepo.FindTop5ByUserId(userId)

	if err != nil {
		return response.DashboardResponse{}, err
	}
	var walletResponse []response.WalletResponse
	for _, item := range wallet {
		walletResponse = append(walletResponse, response.WalletResponse{
			WalletId:     item.WalletId,
			WalletName:   item.WalletName,
			WalletDetail: item.WalletDetail,
			Balance:      item.Balance,
		})
	}

	response := response.DashboardResponse{
		TotalBalance:       totalBalance,
		Wallets:            walletResponse,
		RecentTransactions: transactions,
	}

	return response, nil
}

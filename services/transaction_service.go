package services

import (
	"backend-go-mysql/dto/request"
	"backend-go-mysql/dto/response"
	"backend-go-mysql/models"
	"backend-go-mysql/repositories"
	"backend-go-mysql/utils"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionService struct {
	TransactionRepo repositories.TransactionRepository
	WalletRepo      repositories.WalletRepository
}

func NewTransactionService(transactionRepo repositories.TransactionRepository, walletRepo repositories.WalletRepository) TransactionService {
	return TransactionService{TransactionRepo: transactionRepo,
		WalletRepo: walletRepo}
}

func (s *TransactionService) CreateTransaction(transactionListDto []request.TransactionDto, userId string) ([]response.TransactionResponse, error) {

	var transactionModels []models.Transaction

	var walletId string

	walletBalance := decimal.Zero

	for _, transaction := range transactionListDto {
		walletId = transaction.WalletId
		parsedDate, err := time.Parse("2006-01-02", transaction.TransactionDate)
		if err != nil {
			return []response.TransactionResponse{}, err
		}

		transactionModels = append(transactionModels, models.Transaction{
			TransactionId:   uuid.New().String(),
			WalletId:        transaction.WalletId,
			Name:            transaction.Name,
			Amount:          transaction.Amount,
			Note:            transaction.Note,
			Type:            transaction.Type,
			TransactionDate: parsedDate,
			CategoryId:      transaction.CategoryId,
			UserId:          userId,
			CreatedAt:       utils.NowUTC(),
		})
	}

	err := s.TransactionRepo.CreateTransaction(transactionModels)
	if err != nil {
		return nil, err
	}

	for _, t := range transactionModels {

		switch t.Type {
		case "income":
			walletBalance = walletBalance.Add(t.Amount)

		case "expense":
			walletBalance = walletBalance.Sub(t.Amount)
		}
	}
	s.WalletRepo.UpdateBalance(walletId, walletBalance)

	var responses []response.TransactionResponse
	for _, t := range transactionModels {
		responses = append(responses, response.TransactionResponse{
			TransactionId: t.TransactionId,
			Name:          t.Name,
			Amount:        t.Amount,
			Type:          t.Type,
		})
	}

	return responses, nil
}

func (s *TransactionService) GetAllByWalletId(walletId string, month int, year int) ([]response.TransactionResponse, error) {

	conditionDate := fmt.Sprintf("%04d-%02d", year, month)

	startDateStr := conditionDate + "-01"

	dateFrom, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, err
	}

	dateTo := dateFrom.AddDate(0, 1, 0)

	transactions, err := s.TransactionRepo.FindByWalletId(walletId, dateFrom, dateTo)
	if err != nil {
		return nil, err
	}

	var result []response.TransactionResponse
	for _, t := range transactions {
		category := response.CategoryResponse{
			CategoryId: t.CategoryId,
			Name:       t.Name,
		}

		result = append(result, response.TransactionResponse{
			TransactionId:   t.TransactionId,
			WalletId:        t.WalletId,
			Name:            t.Name,
			Amount:          t.Amount,
			Note:            t.Note,
			Type:            t.Type,
			TransactionDate: t.TransactionDate.String(),
			Category:        category,
			UserId:          t.UserId,
		})
	}

	return result, nil
}

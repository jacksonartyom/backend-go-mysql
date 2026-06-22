package services

import (
	"backend-go-mysql/dto/request"
	"backend-go-mysql/dto/response"
	"backend-go-mysql/models"
	"backend-go-mysql/repositories"
	"backend-go-mysql/utils"
	"errors"

	"github.com/google/uuid"
)

type WalletService struct {
	WalletRepo repositories.WalletRepository
}

func NewWalletService(walletRepo repositories.WalletRepository) WalletService {
	return WalletService{WalletRepo: walletRepo}
}

func (s *WalletService) GetWalletByUserId(userId string) ([]response.WalletResponse, error) {

	wallet, err := s.WalletRepo.FindWalletByUserId(userId)

	if err != nil {
		return []response.WalletResponse{}, errors.New("Data not found")
	}

	var walletResponse []response.WalletResponse

	for _, wallet := range wallet {
		walletResponse = append(walletResponse, response.WalletResponse{
			WalletId:     wallet.WalletId,
			WalletName:   wallet.WalletName,
			WalletDetail: wallet.WalletDetail,
			Balance:      wallet.Balance,
			UserId:       wallet.UserId,
		})
	}

	return walletResponse, nil
}

func (s *WalletService) CreateWallet(walletDto request.WalletDto, userId string) (response.WalletResponse, error) {
	walletRequest := models.Wallet{
		WalletId:     uuid.New().String(),
		WalletName:   walletDto.WalletName,
		WalletDetail: utils.StringPtr(walletDto.WalletDetail),
		Balance:      walletDto.Balance,
		UserId:       userId,
		CreatedAt:    utils.NowUTC(),
	}

	wallet, err := s.WalletRepo.CreateWallet(walletRequest)

	if err != nil {
		return response.WalletResponse{}, errors.New("Data can't save")
	}

	walletResponse := response.WalletResponse{
		WalletId:     wallet.WalletId,
		WalletName:   wallet.WalletName,
		WalletDetail: wallet.WalletDetail,
		Balance:      wallet.Balance,
		UserId:       wallet.UserId,
	}
	return walletResponse, nil
}

func (s *WalletService) UpdateWallet(walletId string, walletDto request.WalletDto) (response.WalletResponse, error) {
	walletRequest := models.Wallet{
		WalletName:   walletDto.WalletName,
		WalletDetail: utils.StringPtr(walletDto.WalletDetail),
		UpdatedAt:    utils.NowUTC(),
	}

	err := s.WalletRepo.UpdateWallet(walletId, &walletRequest)

	if err != nil {
		return response.WalletResponse{}, errors.New("Data can't save")
	}

	walletResponse := response.WalletResponse{
		WalletId:     walletId,
		WalletName:   walletRequest.WalletName,
		WalletDetail: walletRequest.WalletDetail,
	}
	return walletResponse, nil
}

func (s *WalletService) DeletWalletByWalletId(walletId string) error {
	err := s.WalletRepo.DeleteWalletByWalletId(walletId)
	if err != nil {
		return errors.New("Data can't delete")
	}
	return nil
}

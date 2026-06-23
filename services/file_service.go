package services

import (
	"backend-go-mysql/dto/request"
	"io"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
)

type FileService struct {
	transactionService TransactionService
}

func NewFileService(ts TransactionService) FileService {
	return FileService{transactionService: ts}
}

func (s *FileService) ReadExcel(file io.Reader) ([]request.TransactionDto, error) {

	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}

	sheetName := f.GetSheetName(0)

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	var results []request.TransactionDto
	var userId string
	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) < 6 {
			continue
		}

		amountStr := strings.ReplaceAll(row[2], ",", "")

		amount, err := decimal.NewFromString(amountStr)
		if err != nil {
			continue
		}

		userId = row[5]

		item := request.TransactionDto{
			Name:            row[0],
			Note:            row[1],
			Amount:          amount,
			Type:            row[3],
			TransactionDate: row[4],
			UserId:          row[5],
			WalletId:        row[6],
			CategoryId:      row[7],
		}

		results = append(results, item)
	}

	s.transactionService.CreateTransaction(results, userId)

	return results, nil
}

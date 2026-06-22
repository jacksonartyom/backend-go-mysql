package controllers

import (
	"backend-go-mysql/dto/request"
	"backend-go-mysql/dto/response"
	"backend-go-mysql/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService services.TransactionService
}

func (a *TransactionController) CreateTransaction(c *gin.Context) {
	var input []request.TransactionDto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userIdRaw, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	userId, ok := userIdRaw.(string)
	if !ok {
		c.JSON(500, gin.H{"error": "invalid userId type"})
		return
	}

	user, err := a.TransactionService.CreateTransaction(input, userId)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	res := response.SuccessResponse[[]response.TransactionResponse]{
		Message: "success",
		Result:  user,
	}

	c.JSON(200, res)
}

func (a *TransactionController) GetAllTransactionByWalletId(c *gin.Context) {
	// ✅ query params
	walletId := c.Query("wallet_id")
	monthStr := c.Query("month")
	yearStr := c.Query("year")

	// ✅ validate + convert month
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid month",
		})
		return
	}

	// ✅ validate + convert year
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid year",
		})
		return
	}

	// ✅ call service
	transactions, err := a.TransactionService.GetAllByWalletId(walletId, month, year)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	// ✅ response format เหมือน CreateTransaction
	res := response.SuccessResponse[[]response.TransactionResponse]{
		Message: "success",
		Result:  transactions,
	}

	c.JSON(http.StatusOK, res)
}

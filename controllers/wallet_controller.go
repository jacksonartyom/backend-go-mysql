package controllers

import (
	"backend-go-mysql/dto/request"
	"backend-go-mysql/dto/response"
	"backend-go-mysql/services"

	"github.com/gin-gonic/gin"
)

type WalletController struct {
	WalletService services.WalletService
}

func (a *WalletController) GetWallet(c *gin.Context) {

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

	user, err := a.WalletService.GetWalletByUserId(userId)

	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}

	res := response.SuccessResponse[response.WalletResponse]{
		Message: "success",
		Result:  user,
	}

	c.JSON(200, res)
}

func (a *WalletController) CreateWallet(c *gin.Context) {
	var input request.WalletDto

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

	user, err := a.WalletService.CreateWallet(input, userId)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	res := response.SuccessResponse[response.WalletResponse]{
		Message: "success",
		Result:  user,
	}

	c.JSON(200, res)
}

func (a *WalletController) UpdateWallet(c *gin.Context) {
	walletId := c.Param("walletId")

	var input request.WalletDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := a.WalletService.UpdateWallet(walletId, input)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	res := response.SuccessResponse[response.WalletResponse]{
		Message: "success",
		Result:  user,
	}

	c.JSON(200, res)
}

func (a *WalletController) DeleteWallet(c *gin.Context) {

	walletId := c.Param("walletId")

	err := a.WalletService.DeletWalletByWalletId(walletId)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	res := response.SuccessResponse[string]{
		Message: "success",
		Result:  "Delete success",
	}

	c.JSON(200, res)
}

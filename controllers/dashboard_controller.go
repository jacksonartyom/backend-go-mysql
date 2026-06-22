package controllers

import (
	"backend-go-mysql/dto/response"
	"backend-go-mysql/services"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	DashboardService services.DashboardService
}

func (a *DashboardController) GetDashboard(c *gin.Context) {
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

	dashboard, err := a.DashboardService.GetCategoryByUserId(userId)

	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}

	res := response.SuccessResponse[response.DashboardResponse]{
		Message: "success",
		Result:  dashboard,
	}

	c.JSON(200, res)
}

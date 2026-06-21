package controllers

import (
	"backend-go-mysql/dto/request"
	"backend-go-mysql/dto/response"
	"backend-go-mysql/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func (a *UserController) CreateUser(c *gin.Context) {
	var input request.UserDto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := a.UserService.CreateUser(input)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	res := response.SuccessResponse[response.UserResponse]{
		Message: "success",
		Result:  user,
	}

	c.JSON(200, res)
}

func (a *UserController) GetProfile(c *gin.Context) {

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

	user, err := a.UserService.GetProfileByUserId(userId)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	res := response.SuccessResponse[response.UserResponse]{
		Message: "success",
		Result:  user,
	}

	c.JSON(200, res)
}

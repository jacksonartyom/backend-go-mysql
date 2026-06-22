package controllers

import (
	"backend-go-mysql/dto/request"
	"backend-go-mysql/dto/response"
	"backend-go-mysql/services"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryService services.CategoryService
}

func (a *CategoryController) GetCategory(c *gin.Context) {

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

	user, err := a.CategoryService.GetCategoryByUserId(userId)

	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}

	res := response.SuccessResponse[[]response.CategoryResponse]{
		Message: "success",
		Result:  user,
	}

	c.JSON(200, res)
}

func (a *CategoryController) CreateCategory(c *gin.Context) {
	var input request.CategoryDto

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

	input.UserId = userId

	user, err := a.CategoryService.CreateCategory(input)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	res := response.SuccessResponse[response.CategoryResponse]{
		Message: "success",
		Result:  user,
	}

	c.JSON(200, res)
}

func (a *CategoryController) DeleteCategory(c *gin.Context) {

	categoryId := c.Param("categoryId")

	err := a.CategoryService.DeleteCategory(categoryId)

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

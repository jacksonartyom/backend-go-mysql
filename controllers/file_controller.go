package controllers

import (
	"backend-go-mysql/dto/request"
	"backend-go-mysql/dto/response"
	"backend-go-mysql/services"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	FileService services.FileService
}

func (a *FileController) UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file is required"})
		return
	}

	fileName := fmt.Sprintf("%d_%s", time.Now().UnixMilli(), file.Filename)

	uploadDir := "./uploads/"
	filePath := uploadDir + fileName

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(500, gin.H{"error": "cannot create directory"})
		return
	}

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(500, gin.H{"error": "upload failed"})
		return
	}

	fileUrl := "http://localhost:8080/files/" + fileName

	c.JSON(200, gin.H{
		"url": fileUrl,
	})
}

func (a *FileController) UploadExcel(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{"error": "file is required"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot open file"})
		return
	}
	defer f.Close()

	result, err := a.FileService.ReadExcel(f)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res := response.SuccessResponse[[]request.TransactionDto]{
		Message: "success",
		Result:  result,
	}

	c.JSON(200, res)
}

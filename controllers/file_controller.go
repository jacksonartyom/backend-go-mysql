package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type FileController struct {
}

func (a *FileController) UploadFile(c *gin.Context) {
	// 1. รับไฟล์จาก form-data (key = file)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file is required"})
		return
	}

	// 2. ตั้งชื่อไฟล์ใหม่
	fileName := fmt.Sprintf("%d_%s", time.Now().UnixMilli(), file.Filename)

	// 3. path ที่จะเก็บ
	uploadDir := "./uploads/"
	filePath := uploadDir + fileName

	// 4. สร้าง folder (ถ้ายังไม่มี)
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(500, gin.H{"error": "cannot create directory"})
		return
	}

	// 5. save file
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(500, gin.H{"error": "upload failed"})
		return
	}

	// 6. สร้าง URL
	fileUrl := "http://localhost:8080/files/" + fileName

	c.JSON(200, gin.H{
		"url": fileUrl,
	})
}

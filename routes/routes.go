package routes

import (
	"backend-go-mysql/internal/container"
	"backend-go-mysql/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, app container.App) {

	v1 := r.Group("/api/v1")

	v1.POST("/sign-in", app.AuthController.Login)
	v1.POST("/sign-up", app.UserController.CreateUser)

	auth := r.Group("/api/v1")
	auth.Use(middleware.AuthMiddleware(app.JwtService))

	auth.GET("/user/user-profile", app.UserController.GetProfile)

	// upload file (protected)
	auth.POST("/upload/profile", app.FileController.UploadFile)

	// serve static file (ต้องอยู่นอก group)
	r.Static("/files", "./uploads")
	// ถ้ามีเพิ่ม
	auth.GET("/wallet", app.WalletController.GetWallet)
	auth.POST("/wallet", app.WalletController.CreateWallet)
	auth.PUT("/wallet/:walletId", app.WalletController.UpdateWallet)
	auth.DELETE("/wallet/:walletId", app.WalletController.DeleteWallet)
}

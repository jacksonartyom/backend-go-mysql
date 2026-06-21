package container

import (
	"backend-go-mysql/controllers"
	"backend-go-mysql/repositories"
	"backend-go-mysql/services"
	"backend-go-mysql/utils"
)

type App struct {
	AuthController   controllers.AuthController
	UserController   controllers.UserController
	FileController   controllers.FileController
	JwtService       *utils.JwtService
	WalletController controllers.WalletController
}

func Build(jwtService *utils.JwtService) App {
	userRepo := repositories.NewUserRepository()
	walletRepo := repositories.NewWalletRepository()

	authService := services.NewAuthService(userRepo, jwtService)
	userService := services.NewUserService(userRepo)
	walletService := services.NewWalletService(walletRepo)

	return App{
		AuthController:   controllers.AuthController{AuthService: authService},
		UserController:   controllers.UserController{UserService: userService},
		JwtService:       jwtService,
		WalletController: controllers.WalletController{WalletService: walletService},
	}
}

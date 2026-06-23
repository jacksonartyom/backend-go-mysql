package container

import (
	"backend-go-mysql/controllers"
	"backend-go-mysql/repositories"
	"backend-go-mysql/services"
	"backend-go-mysql/utils"
)

type App struct {
	AuthController        controllers.AuthController
	UserController        controllers.UserController
	FileController        controllers.FileController
	JwtService            *utils.JwtService
	WalletController      controllers.WalletController
	CategoryController    controllers.CategoryController
	TransactionController controllers.TransactionController
	DashboardController   controllers.DashboardController
}

func Build(jwtService *utils.JwtService) App {
	userRepo := repositories.NewUserRepository()
	walletRepo := repositories.NewWalletRepository()
	categoryRepo := repositories.NewCategoryRepository()
	transactionRepo := repositories.NewTransactionRepository()

	authService := services.NewAuthService(userRepo, jwtService)
	userService := services.NewUserService(userRepo)
	walletService := services.NewWalletService(walletRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	transactionService := services.NewTransactionService(transactionRepo, walletRepo)
	dashboardService := services.NewDashboardService(transactionRepo, walletRepo)
	fileService := services.NewFileService(transactionService)

	return App{
		AuthController:        controllers.AuthController{AuthService: authService},
		UserController:        controllers.UserController{UserService: userService},
		JwtService:            jwtService,
		WalletController:      controllers.WalletController{WalletService: walletService},
		CategoryController:    controllers.CategoryController{CategoryService: categoryService},
		TransactionController: controllers.TransactionController{TransactionService: transactionService},
		DashboardController:   controllers.DashboardController{DashboardService: dashboardService},
		FileController:        controllers.FileController{FileService: fileService},
	}
}

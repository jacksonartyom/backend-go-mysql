package repositories

import (
	"backend-go-mysql/config"
	"backend-go-mysql/models"
)

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (r *UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)

	return user, result.Error
}

func (r *UserRepository) CreateUser(user models.User) (models.User, error) {
	result := config.DB.Create(&user)

	return user, result.Error
}

func (r *UserRepository) FindByUserId(userId string) (models.User, error) {
	var user models.User
	result := config.DB.Where("user_id = ?", userId).First(&user)

	return user, result.Error
}

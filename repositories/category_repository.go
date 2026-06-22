package repositories

import (
	"backend-go-mysql/config"
	"backend-go-mysql/models"
)

type CategoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{}
}

func (r *CategoryRepository) FindByUserId(userId string) ([]models.Category, error) {
	var Category []models.Category
	result := config.DB.Where("is_system_default = true or user_id = ?", userId).Find(&Category)

	return Category, result.Error
}

func (r *CategoryRepository) CreateCategory(Category models.Category) (models.Category, error) {
	result := config.DB.Create(&Category)

	return Category, result.Error
}

func (r *CategoryRepository) DeleteCategoryByCategoryId(categoryId string) error {
	return config.DB.
		Where("category_id = ?", categoryId).
		Delete(&models.Category{}).Error
}

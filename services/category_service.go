package services

import (
	"backend-go-mysql/dto/request"
	"backend-go-mysql/dto/response"
	"backend-go-mysql/models"
	"backend-go-mysql/repositories"
	"backend-go-mysql/utils"
	"errors"

	"github.com/google/uuid"
)

type CategoryService struct {
	CategoryRepo repositories.CategoryRepository
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryService {
	return CategoryService{CategoryRepo: categoryRepo}
}

func (s *CategoryService) GetCategoryByUserId(userId string) ([]response.CategoryResponse, error) {
	categories, err := s.CategoryRepo.FindByUserId(userId)

	if err != nil {
		return nil, errors.New("Data not found")
	}

	var categoryResponses []response.CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses, response.CategoryResponse{
			CategoryId: category.CategoryId,
			Name:       category.Name,
			Type:       category.Type,
		})
	}

	return categoryResponses, nil
}

func (s *CategoryService) CreateCategory(categoryDto request.CategoryDto) (response.CategoryResponse, error) {
	categoryRequest := models.Category{
		CategoryId:      uuid.New().String(),
		Name:            categoryDto.Name,
		Type:            categoryDto.Type,
		UserId:          categoryDto.UserId,
		IsSystemDefault: false,
		CreatedAt:       utils.NowUTC(),
	}
	category, err := s.CategoryRepo.CreateCategory(categoryRequest)

	if err != nil {
		return response.CategoryResponse{}, errors.New("Data can't save")
	}

	categoryResponse := response.CategoryResponse{
		CategoryId: category.CategoryId,
		Name:       category.Name,
		Type:       category.Type,
	}

	return categoryResponse, nil
}

func (s *CategoryService) DeleteCategory(categoryId string) error {
	err := s.CategoryRepo.DeleteCategoryByCategoryId(categoryId)
	if err != nil {
		return errors.New("Data can't delete")
	}
	return nil
}

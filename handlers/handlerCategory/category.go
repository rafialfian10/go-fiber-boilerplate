package handlerCategory

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/repositories"
)

type handlerCategory struct {
	CategoryRepository repositories.CategoryRepository
}

func HandlerCategory(categoryRepository repositories.CategoryRepository) *handlerCategory {
	return &handlerCategory{categoryRepository}
}

func convertCategoryResponse(category *models.Category) *dto.CategoryResponse {
	return &dto.CategoryResponse{
		ID:       category.ID,
		Category: category.Category,
	}
}

func convertMultipleCategoryResponse(categoryDatas *[]models.Category) *[]dto.CategoryResponse {
	var categories []dto.CategoryResponse

	for _, c := range *categoryDatas {
		categories = append(categories, *convertCategoryResponse(&c))
	}

	return &categories
}

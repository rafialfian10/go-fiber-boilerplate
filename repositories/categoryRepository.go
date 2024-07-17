package repositories

import (
	"fmt"
	"go-restapi-boilerplate/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategories(limit, offset int, searchQuery string) (*[]models.Category, int64, error)
	GetCategoryByID(categoryID uint) (*models.Category, error)
	CreateCategory(category *models.Category) (*models.Category, error)
	UpdateCategory(category *models.Category) (*models.Category, error)
	DeleteCategory(category *models.Category) (*models.Category, error)
}

func (r *repository) GetCategories(limit, offset int, searchQuery string) (*[]models.Category, int64, error) {
	var (
		categories    []models.Category
		totalCategory int64
	)

	trx := r.db.Session(&gorm.Session{})

	if searchQuery != "" {
		trx = trx.Where("category LIKE ?", fmt.Sprintf("%%%s%%", searchQuery))
	}

	trx.Model(&models.Category{}).Count(&totalCategory)

	err := trx.Limit(limit).Offset(offset).Find(&categories).Error

	return &categories, totalCategory, err
}

func (r *repository) GetCategoryByID(categoryID uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ?", categoryID).First(&category).Error

	return &category, err
}

func (r *repository) CreateCategory(category *models.Category) (*models.Category, error) {
	err := r.db.Create(category).Error

	return category, err
}

func (r *repository) UpdateCategory(category *models.Category) (*models.Category, error) {
	err := r.db.Model(category).Updates(*category).Error

	return category, err
}

func (r *repository) DeleteCategory(category *models.Category) (*models.Category, error) {
	err := r.db.Delete(category).Error

	return category, err
}

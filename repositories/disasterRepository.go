package repositories

import (
	"fmt"
	"go-restapi-boilerplate/models"

	"gorm.io/gorm"
)

type DisasterRepository interface {
	GetDisasters(limit, offset int, searchQuery string) (*[]models.Disaster, int64, error)
	GetDisasterByID(disasterID uint) (*models.Disaster, error)
	CreateDisaster(disaster *models.Disaster) (*models.Disaster, error)
	UpdateDisaster(disaster *models.Disaster) (*models.Disaster, error)
	DeleteDisaster(disaster *models.Disaster) (*models.Disaster, error)
}

func (r *repository) GetDisasters(limit, offset int, searchQuery string) (*[]models.Disaster, int64, error) {
	var (
		disasters     []models.Disaster
		totalDisaster int64
	)

	trx := r.db.Session(&gorm.Session{})

	if searchQuery != "" {
		trx = trx.Where("disaster LIKE ?", fmt.Sprintf("%%%s%%", searchQuery))
	}

	trx.Model(&models.Disaster{}).Count(&totalDisaster)

	err := trx.Limit(limit).Offset(offset).Preload("User.Role").Preload("Category").Find(&disasters).Error

	return &disasters, totalDisaster, err
}

func (r *repository) GetDisasterByID(disasterID uint) (*models.Disaster, error) {
	var disaster models.Disaster
	err := r.db.Where("id = ?", disasterID).Preload("User.Role").Preload("Category").First(&disaster).Error

	return &disaster, err
}

func (r *repository) CreateDisaster(disaster *models.Disaster) (*models.Disaster, error) {
	err := r.db.Preload("User.Role").Preload("Category").Create(disaster).Error

	return disaster, err
}

func (r *repository) UpdateDisaster(disaster *models.Disaster) (*models.Disaster, error) {
	// err := r.db.Preload("User.Role").Preload("Category").Save(disaster).Error
	err := r.db.Model(&models.Disaster{}).Where("id = ?", disaster.ID).Updates(map[string]interface{}{
		"Title":        disaster.Title,
		"Description":  disaster.Description,
		"Location":     disaster.Location,
		"CategoryID":   disaster.CategoryID,
		"Date":         disaster.Date,
		"Donate":       disaster.Donate,
		"DonateTarget": disaster.DonateTarget,
		"Image":        disaster.Image,
		"IsTrending":   disaster.IsTrending,
	}).Error

	return disaster, err
}

func (r *repository) DeleteDisaster(disaster *models.Disaster) (*models.Disaster, error) {
	err := r.db.Preload("User.Role").Preload("Category").Delete(disaster).Error

	return disaster, err
}

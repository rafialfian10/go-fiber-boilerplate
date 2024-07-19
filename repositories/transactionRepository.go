package repositories

import (
	"fmt"
	"go-restapi-boilerplate/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransactionsByAdmin(limit, offset int, searchQuery string) (*[]models.Transaction, int64, error)
	GetTransactionsByUser(userID uuid.UUID, limit, offset int, searchQuery string) (*[]models.Transaction, int64, error)
	GetTransactionByID(transactionID uint) (*models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) (*models.Transaction, error)
	UpdateTransaction(transaction *models.Transaction) (*models.Transaction, error)
	UpdateTokenTransaction(token string, transactionID uint) (*models.Transaction, error)
	DeleteTransaction(transaction *models.Transaction) (*models.Transaction, error)
}

func (r *repository) GetTransactionsByAdmin(limit, offset int, searchQuery string) (*[]models.Transaction, int64, error) {
	var (
		transactions     []models.Transaction
		totalTransaction int64
	)

	trx := r.db.Session(&gorm.Session{})

	if searchQuery != "" {
		trx = trx.Where("transaction LIKE ?", fmt.Sprintf("%%%s%%", searchQuery))
	}

	trx.Model(&models.Transaction{}).Count(&totalTransaction)

	err := trx.Limit(limit).Offset(offset).Preload("User.Role").Preload("Disaster.User.Role").Preload("Disaster.Category").Find(&transactions).Error

	return &transactions, totalTransaction, err
}

func (r *repository) GetTransactionsByUser(userID uuid.UUID, limit, offset int, searchQuery string) (*[]models.Transaction, int64, error) {
	var (
		transactions     []models.Transaction
		totalTransaction int64
	)

	trx := r.db.Session(&gorm.Session{})

	if searchQuery != "" {
		trx = trx.Where("transaction LIKE ?", fmt.Sprintf("%%%s%%", searchQuery))
	}

	// Filter transactions by user ID
	trx = trx.Where("user_id = ?", userID)

	trx.Model(&models.Transaction{}).Count(&totalTransaction)

	err := trx.Limit(limit).Offset(offset).Preload("User.Role").Preload("Disaster.User.Role").Preload("Disaster.Category").Find(&transactions).Error

	return &transactions, totalTransaction, err
}

func (r *repository) GetTransactionByID(transactionID uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Where("id = ?", transactionID).Preload("User.Role").Preload("Disaster.User.Role").Preload("Disaster.Category").First(&transaction).Error

	return &transaction, err
}

func (r *repository) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	err := r.db.Preload("User.Role").Preload("Disaster.User.Role").Preload("Disaster.Category").Create(transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	err := r.db.Preload("User.Role").Preload("Disaster.User.Role").Preload("Disaster.Category").Save(transaction).Error

	return transaction, err
}

func (r *repository) UpdateTokenTransaction(token string, transactionID uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User.Role").Preload("Disaster.User.Role").Preload("Disaster.Category").First(&transaction, "id = ?", transactionID).Error
	if err != nil {
		return nil, err
	}

	transaction.Token = token
	err = r.db.Model(&transaction).Updates(transaction).Error

	return &transaction, err
}

func (r *repository) DeleteTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	err := r.db.Preload("User.Role").Preload("Disaster.User.Role").Preload("Disaster.Category").Delete(transaction).Error

	return transaction, err
}

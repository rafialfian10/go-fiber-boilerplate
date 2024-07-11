package repositories

import (
	"fmt"
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(limit, offset int, filter dto.UserFilter, searchQuery string) (*[]models.Users, int64, error)
	GetUserByID(id uuid.UUID) (*models.Users, error)
	CreateUser(user *models.Users) (*models.Users, error)
	UpdateUser(user *models.Users) (*models.Users, error)
	UpdateUserByAdmin(user *models.Users) (*models.Users, error)
	DeleteUser(user *models.Users, id uuid.UUID) (*models.Users, error)
	GetUserByEmail(email string) (*models.Users, error)
}

func (r *repository) GetUsers(limit, offset int, filter dto.UserFilter, searchQuery string) (*[]models.Users, int64, error) {
	var (
		users     []models.Users
		totalUser int64
	)

	// create new transaction
	trx := r.db.Session(&gorm.Session{})

	if filter.RoleID != 0 {
		trx = trx.Where("role_id = ?", filter.RoleID)
	}

	// join tables, used for complex searching on relation table
	trx = trx.Joins("JOIN roles ON roles.id = users.role_id")

	if searchQuery != "" {
		searchQuery = fmt.Sprintf("%%%s%%", searchQuery)

		trx = trx.Where("full_name LIKE ? OR email LIKE ? OR phone LIKE ? OR address LIKE ? OR roles.role LIKE ?",
			searchQuery, // full_name
			searchQuery, // email
			searchQuery, // phone
			searchQuery, // gender
			searchQuery, // address
			searchQuery) // role
	}

	// preloading, used for get relation data for results
	trx = trx.Preload("Role")

	// count transaction result
	trx.Model(&models.Users{}).
		Count(&totalUser)

	// set pagination
	err := trx.Limit(limit).
		Offset(offset).
		Find(&users).Error

	return &users, totalUser, err
}

func (r *repository) GetUserByID(id uuid.UUID) (*models.Users, error) {
	var user models.Users

	err := r.db.Preload("Role").Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *repository) CreateUser(user *models.Users) (*models.Users, error) {
	err := r.db.Create(user).Error

	return user, err
}

func (r *repository) UpdateUser(user *models.Users) (*models.Users, error) {
	err := r.db.Preload("Role").Save(user).Error
	return user, err
}

func (r *repository) UpdateUserByAdmin(user *models.Users) (*models.Users, error) {
	query := fmt.Sprintf(`update users set full_name = '%s', email = '%s', is_email_verified = '%t', phone = '%s', is_phone_verified = '%t', gender = '%s', address = '%s', password = '%s', role_id = '%d', image = '%s' where id = '%s'`, user.FullName, user.Email, user.IsEmailVerified, user.Phone, user.IsPhoneVerified, user.Gender, user.Address, user.Password, user.RoleID, user.Image, user.ID)

	err := r.db.Exec(query).Error

	return user, err
}

func (r *repository) DeleteUser(user *models.Users, id uuid.UUID) (*models.Users, error) {
	err := r.db.Where("id = ?", id).Delete(user).Error
	return user, err
}

func (r *repository) GetUserByEmail(email string) (*models.Users, error) {
	var user models.Users

	err := r.db.Preload("Role").Where("email = ?", email).First(&user).Error
	return &user, err
}

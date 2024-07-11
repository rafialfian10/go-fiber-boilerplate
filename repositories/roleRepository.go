package repositories

import (
	"fmt"
	"go-restapi-boilerplate/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRoles(limit, offset int, searchQuery string) (*[]models.Role, int64, error)
	GetRoleByID(roleID uint) (*models.Role, error)
	CreateRole(role *models.Role) (*models.Role, error)
	UpdateRole(role *models.Role) (*models.Role, error)
	DeleteRole(role *models.Role) (*models.Role, error)
	CheckIsRoleUsed(role *models.Role) (bool, error)
}

func (r *repository) GetRoles(limit, offset int, searchQuery string) (*[]models.Role, int64, error) {
	var (
		roles     []models.Role
		totalRole int64
	)

	trx := r.db.Session(&gorm.Session{})

	if searchQuery != "" {
		trx = trx.Where("role LIKE ?", fmt.Sprintf("%%%s%%", searchQuery))
	}

	trx.Model(&models.Role{}).Count(&totalRole)

	err := trx.Limit(limit).Offset(offset).Find(&roles).Error

	return &roles, totalRole, err
}

func (r *repository) GetRoleByID(roleID uint) (*models.Role, error) {
	var role models.Role
	err := r.db.Where("id = ?", roleID).First(&role).Error

	return &role, err
}

func (r *repository) CreateRole(role *models.Role) (*models.Role, error) {
	err := r.db.Create(role).Error

	return role, err
}

func (r *repository) UpdateRole(role *models.Role) (*models.Role, error) {
	err := r.db.Model(role).Updates(*role).Error

	return role, err
}

func (r *repository) DeleteRole(role *models.Role) (*models.Role, error) {
	err := r.db.Delete(role).Error

	return role, err
}

func (r *repository) CheckIsRoleUsed(role *models.Role) (bool, error) {
	var count int

	err := r.db.Raw("select count(*) from roles mr join users mu on mr.id  = mu.role_id where mr.id = ?", role.ID).Scan(&count).Error

	if count > 0 {
		return true, err
	}

	return false, err
}

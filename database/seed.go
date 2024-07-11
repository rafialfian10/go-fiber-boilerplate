package database

import (
	"errors"
	"fmt"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/pkg/bcrypt"
	"go-restapi-boilerplate/pkg/mysql"
	"log"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RunSeeder() {
	// Role
	if mysql.DB.Migrator().HasTable(&models.Role{}) {
		newRole := []models.Role{}

		newRole = append(newRole, models.Role{
			Role: "Superadmin",
		})
		newRole = append(newRole, models.Role{
			Role: "Admin",
		})
		newRole = append(newRole, models.Role{
			Role: "User",
		})

		for _, role := range newRole {
			errAddRole := mysql.DB.Create(&role).Error
			if errAddRole != nil {
				fmt.Println(errAddRole.Error())
				log.Fatal("Seeding failed")
			}
		}

		fmt.Println("Success seeding master role...")
	}

	// Add Superadmin
	if mysql.DB.Migrator().HasTable(&models.Users{}) {
		// check is user table has minimum 1 user
		err := mysql.DB.First(&models.Users{}).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// create 1 user
			newUser := models.Users{
				ID:              uuid.New(),
				FullName:        "Super Admin",
				Email:           os.Getenv("SUPERADMIN_EMAIL"),
				IsEmailVerified: true,
				IsPhoneVerified: true,
				RoleID:          1,
			}

			hashPassword, err := bcrypt.HashingPassword(os.Getenv("SUPERADMIN_PASSWORD"))
			if err != nil {
				log.Fatal("Hash password failed")
			}

			newUser.Password = hashPassword

			// insert user to database
			errAddUser := mysql.DB.Create(&newUser).Error
			if errAddUser != nil {
				fmt.Println(errAddUser.Error())
				log.Fatal("Seeding failed")
			}
		}
		fmt.Println("Success seeding super admin...")
	}

	fmt.Println("Seeding completed successfully")
}

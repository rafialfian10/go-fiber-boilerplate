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
		fmt.Println("Roles table exists. Seeding roles...")

		newRole := []models.Role{
			{Role: "Superadmin"},
			{Role: "Admin"},
			{Role: "User"},
		}

		for _, role := range newRole {
			errAddRole := mysql.DB.FirstOrCreate(&role, models.Role{Role: role.Role}).Error
			if errAddRole != nil {
				fmt.Println("Error seeding role:", errAddRole)
				log.Fatal("Seeding failed")
			}
		}

		fmt.Println("Successfully seeded roles")
	} else {
		log.Fatal("Table 'roles' doesn't exist")
	}

	// Add Superadmin and Admin
	if mysql.DB.Migrator().HasTable(&models.Users{}) {
		fmt.Println("Users table exists. Seeding superadmin and admin users...")

		// Check if user table has minimum 1 user
		var user models.Users
		err := mysql.DB.First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("No users found. Seeding default users...")

			// Create superadmin user
			newSuperAdmin := models.Users{
				ID:              uuid.New(),
				FullName:        "Super Admin",
				Email:           os.Getenv("SUPERADMIN_EMAIL"),
				IsEmailVerified: true,
				Phone:           os.Getenv("SUPERADMIN_PHONE"),
				IsPhoneVerified: true,
				RoleID:          1,
			}

			hashPassword, err := bcrypt.HashingPassword(os.Getenv("SUPERADMIN_PASSWORD"))
			if err != nil {
				log.Fatal("Hash password failed:", err)
			}
			newSuperAdmin.Password = hashPassword

			// Insert superadmin user to database
			errAddUser := mysql.DB.Create(&newSuperAdmin).Error
			if errAddUser != nil {
				fmt.Println("Error seeding superadmin user:", errAddUser)
				log.Fatal("Seeding failed")
			}

			// Create admin user
			newAdmin := models.Users{
				ID:              uuid.New(),
				FullName:        "Admin",
				Email:           os.Getenv("ADMIN_EMAIL"),
				IsEmailVerified: true,
				Phone:           os.Getenv("ADMIN_EMAIL"),
				IsPhoneVerified: true,
				RoleID:          2,
			}

			hashPassword, err = bcrypt.HashingPassword(os.Getenv("ADMIN_PASSWORD"))
			if err != nil {
				log.Fatal("Hash password failed:", err)
			}
			newAdmin.Password = hashPassword

			// Insert admin user to database
			errAddUser = mysql.DB.Create(&newAdmin).Error
			if errAddUser != nil {
				fmt.Println("Error seeding admin user:", errAddUser)
				log.Fatal("Seeding failed")
			}

			fmt.Println("Successfully seeded superadmin and admin users")
		} else if err != nil {
			fmt.Println("Error checking for existing users:", err)
			log.Fatal("Seeding failed")
		} else {
			fmt.Println("Users already exist. Skipping user seeding")
		}
	} else {
		log.Fatal("Table 'users' doesn't exist")
	}

	fmt.Println("Seeding completed successfully")
}

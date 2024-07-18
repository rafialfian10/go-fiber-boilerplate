package database

import (
	"fmt"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/pkg/mysql"
	"log"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Log{},
		&models.Role{},
		&models.User{},
		&models.Category{},
		&models.Disaster{},
		&models.Image{},
		&models.Transaction{},
		// put another models struct here
	)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Migration failed")
	}

	fmt.Println("Migration up completed successfully")
}

func DropMigration() {
	err := mysql.DB.Migrator().DropTable(
		&models.Log{},
		&models.Role{},
		&models.User{},
		&models.Category{},
		&models.Disaster{},
		&models.Image{},
		&models.Transaction{},
		// put another models struct here
	)

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Migration failed")
	}

	fmt.Println("Migration down completed successfully")
}

func IsDatabaseSeeded() bool {
	var count int64
	mysql.DB.Table("users").Count(&count)
	return count > 0
}

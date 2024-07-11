package database

import (
	"fmt"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/pkg/mysql"
	"log"
)

// migration up
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Log{},
		&models.Role{}, &models.Users{},
		// put another models struct here
	)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Migration failed")
	}

	fmt.Println("Migration up completed successfully")
}

// migration down
func DropMigration() {
	err := mysql.DB.Migrator().DropTable(
		&models.Log{},
		&models.Role{}, &models.Users{},
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

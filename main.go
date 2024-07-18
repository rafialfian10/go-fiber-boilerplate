package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/asidikrdn/otptimize"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"go-restapi-boilerplate/database"
	"go-restapi-boilerplate/pkg/middleware"
	"go-restapi-boilerplate/pkg/mysql"
	"go-restapi-boilerplate/routes"
)

func main() {
	// Load environment variables
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading environment variables file:", errEnv)
	}

	// Database initialization
	mysql.DatabaseInit()

	// otptimize connection init
	mailPort, _ := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	mailConfig := otptimize.MailConfig{
		Host:     os.Getenv("CONFIG_SMTP_HOST"),
		Port:     mailPort,
		Email:    os.Getenv("CONFIG_AUTH_EMAIL"),
		Password: os.Getenv("CONFIG_AUTH_PASSWORD"),
	}
	redisConfig := otptimize.RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
	otptimize.ConnectionInit(mailConfig, redisConfig)

	// Database seeder and migration
	database.RunMigration()
	database.RunSeeder()
	// database.DropMigration()

	// Create a new Fiber instance
	app := fiber.New()

	// Middleware
	// app.Use(middleware.UserAuth())
	app.Use(middleware.Logger())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Replace with your allowed origins
		AllowMethods: "HEAD, OPTIONS, GET, POST, PUT, PATCH, DELETE",
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Authorization",
	}))

	// Initialize routes
	routes.RouterInit(app.Group("/api/v1"))

	// Static files
	app.Static("/static", "./uploads")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	fmt.Println("Server running on localhost:" + port)
	log.Fatal(app.Listen(":" + port))
}

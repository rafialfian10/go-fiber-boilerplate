package routes

import (
	"go-restapi-boilerplate/handlers/handlerAuth"
	"go-restapi-boilerplate/pkg/middleware"
	"go-restapi-boilerplate/pkg/mysql"
	"go-restapi-boilerplate/repositories"

	"github.com/gofiber/fiber/v2"
)

func Auth(r fiber.Router) {
	userRepository := repositories.MakeRepository(mysql.DB)
	h := handlerAuth.HandlerAuth(userRepository)

	r.Post("/register", h.RegisterUser)
	r.Post("/login", h.Login)
	r.Get("/check-auth", middleware.UserAuth(), h.CheckAuth)
	r.Get("/resend-otp/:email", h.ResendOTP)
	r.Post("/verify-email", h.VerifyEmail)
}

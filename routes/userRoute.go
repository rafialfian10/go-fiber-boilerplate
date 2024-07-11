package routes

import (
	"go-restapi-boilerplate/handlers/handlerUser"
	"go-restapi-boilerplate/pkg/middleware"
	"go-restapi-boilerplate/pkg/mysql"
	"go-restapi-boilerplate/repositories"

	"github.com/gofiber/fiber/v2"
)

func User(r fiber.Router) {
	userRepository := repositories.MakeRepository(mysql.DB)
	h := handlerUser.HandlerUser(userRepository)

	r.Get("/users", middleware.AdminAuth(), h.GetUsers)
	r.Get("/user/profile", middleware.UserAuth(), h.GetProfile)
	r.Get("/user/:id", middleware.UserAuth(), h.GetUserByID)
	r.Patch("/user/:id", middleware.UserAuth(), middleware.UploadSingleFile(), h.UpdateUser)
	r.Patch("/user-by-admin/:id", middleware.AdminAuth(), middleware.UploadSingleFile(), h.UpdateUserByAdmin)
	r.Delete("user/:id", middleware.SuperAdminAuth(), h.DeleteUser)
}

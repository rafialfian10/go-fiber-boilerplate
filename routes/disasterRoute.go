package routes

import (
	"go-restapi-boilerplate/handlers/handlerDisaster"
	"go-restapi-boilerplate/pkg/middleware"
	"go-restapi-boilerplate/pkg/mysql"
	"go-restapi-boilerplate/repositories"

	"github.com/gofiber/fiber/v2"
)

func Disaster(r fiber.Router) {
	disasterRepository := repositories.MakeRepository(mysql.DB)
	h := handlerDisaster.HandlerDisaster(disasterRepository)

	r.Get("/disasters", middleware.UserAuth(), h.GetDisasters)
	r.Get("/disaster/:id", middleware.UserAuth(), h.GetDisasterByID)
	r.Post("/disaster", middleware.AdminAuth(), middleware.UploadSingleImage(), h.CreateDisaster)
	r.Patch("/disaster/:id", middleware.AdminAuth(), middleware.UploadSingleImage(), h.UpdateDisaster)
	r.Delete("/disaster/:id", middleware.AdminAuth(), h.DeleteDisaster)
}

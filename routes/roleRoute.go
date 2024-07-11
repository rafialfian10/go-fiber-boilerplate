package routes

import (
	"go-restapi-boilerplate/handlers/handlerRole"
	"go-restapi-boilerplate/pkg/middleware"
	"go-restapi-boilerplate/pkg/mysql"
	"go-restapi-boilerplate/repositories"

	"github.com/gofiber/fiber/v2"
)

func Role(r fiber.Router) {
	roleRepository := repositories.MakeRepository(mysql.DB)
	h := handlerRole.HandlerRole(roleRepository)

	r.Get("/roles", middleware.AdminAuth(), h.GetRoles)
	r.Get("/role/:id", middleware.AdminAuth(), h.GetRoleByID)
	r.Post("/role", middleware.SuperAdminAuth(), h.CreateRole)
	r.Patch("/role/:id", middleware.SuperAdminAuth(), h.UpdateRole)
	r.Delete("/role/:id", middleware.SuperAdminAuth(), h.DeleteRole)
}

package routes

import (
	"go-restapi-boilerplate/handlers/handlerCategory"
	"go-restapi-boilerplate/pkg/middleware"
	"go-restapi-boilerplate/pkg/mysql"
	"go-restapi-boilerplate/repositories"

	"github.com/gofiber/fiber/v2"
)

func Category(r fiber.Router) {
	categoryRepository := repositories.MakeRepository(mysql.DB)
	h := handlerCategory.HandlerCategory(categoryRepository)

	r.Get("/categories", middleware.UserAuth(), h.GetCategories)
	r.Get("/category/:id", middleware.UserAuth(), h.GetCategoryByID)
	r.Post("/category", middleware.UserAuth(), h.CreateCategory)
	r.Patch("/category/:id", middleware.UserAuth(), h.UpdateCategory)
	r.Delete("/category/:id", middleware.UserAuth(), h.DeleteCategory)
}

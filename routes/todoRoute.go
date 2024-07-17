package routes

import (
	"go-restapi-boilerplate/handlers/handlerTodo"
	"go-restapi-boilerplate/pkg/middleware"
	"go-restapi-boilerplate/pkg/mysql"
	"go-restapi-boilerplate/repositories"

	"github.com/gofiber/fiber/v2"
)

func Todo(r fiber.Router) {
	todoRepository := repositories.MakeRepository(mysql.DB)
	h := handlerTodo.HandlerTodo(todoRepository)

	r.Get("/todos", middleware.UserAuth(), h.GetTodos)
	r.Get("/todo/:id", middleware.UserAuth(), h.GetTodoByID)
	r.Post("/todo", middleware.UserAuth(), h.CreateTodo)
	r.Patch("/todo/:id", middleware.UserAuth(), h.UpdateTodo)
	r.Delete("/todo/:id", middleware.UserAuth(), h.DeleteTodo)
}

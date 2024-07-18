package routes

import (
	"go-restapi-boilerplate/handlers/handlerTransaction"
	"go-restapi-boilerplate/pkg/middleware"
	"go-restapi-boilerplate/pkg/mysql"
	"go-restapi-boilerplate/repositories"

	"github.com/gofiber/fiber/v2"
)

func Transaction(r fiber.Router) {
	transactionRepository := repositories.MakeRepository(mysql.DB)
	h := handlerTransaction.HandlerTransaction(transactionRepository)

	r.Get("/transactions-by-admin", middleware.AdminAuth(), h.GetTransactionsByAdmin)
	r.Get("/transactions-by-user", middleware.UserAuth(), h.GetTransactionsByUser)
	r.Get("/transaction/:id", middleware.UserAuth(), h.GetTransactionByID)
	r.Post("/transaction", middleware.UserAuth(), h.CreateTransaction)
	r.Patch("/transaction/:id", middleware.AdminAuth(), h.UpdateTransaction)
	r.Delete("/transaction/:id", middleware.AdminAuth(), h.DeleteTransaction)
	r.Post("/notification", h.Notification)
}

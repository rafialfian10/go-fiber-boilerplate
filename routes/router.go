package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouterInit(app fiber.Router) {
	User(app)
	Role(app)
	Auth(app)
	Category(app)
	Todo(app)
}

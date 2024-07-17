package handlerTodo

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (h *handlerTodo) CreateTodo(c *fiber.Ctx) error {
	var request dto.CreateTodoRequest

	err := c.BodyParser(&request)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	claims, ok := c.Locals("userData").(jwt.MapClaims)
	if !ok {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User data from jwt payload is not found",
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// Extract user data from JWT claims
	userId, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	date, _ := time.Parse("2006-01-02", c.FormValue("date"))

	todo := models.Todo{
		UserID:      userId,
		Title:       request.Title,
		Description: request.Description,
		CategoryID:  request.CategoryID,
		Date:        date,
	}

	addedTodo, err := h.TodoRepository.CreateTodo(&todo)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	newTodo, err := h.TodoRepository.GetTodoByID(addedTodo.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusCreated,
		Message: "Category successfully created",
		Data:    convertTodoResponse(newTodo),
	}
	return c.Status(http.StatusCreated).JSON(response)
}

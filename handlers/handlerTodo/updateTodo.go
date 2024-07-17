package handlerTodo

import (
	"go-restapi-boilerplate/dto"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerTodo) UpdateTodo(c *fiber.Ctx) error {
	var request dto.UpdateTodoRequest

	if err := c.BodyParser(&request); err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Invalid todo ID",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	todo, err := h.TodoRepository.GetTodoByID(uint(id))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	if request.Title != "" && request.Title != todo.Title {
		todo.Title = request.Title
	}

	if request.Description != "" && request.Description != todo.Description {
		todo.Description = request.Description
	}

	if request.CategoryID != 0 && request.CategoryID != todo.CategoryID {
		todo.CategoryID = request.CategoryID
	}

	if request.Date != "" {
		date, err := time.Parse("2006-01-02", request.Date)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: "Invalid date format",
			}
			return c.Status(http.StatusBadRequest).JSON(response)
		}
		todo.Date = date
	}

	if request.IsDone {
		todo.IsDone = request.IsDone
	}

	updatedTodo, err := h.TodoRepository.UpdateTodo(todo)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	todo, err = h.TodoRepository.GetTodoByID(updatedTodo.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "Todo successfully updated",
		Data:    convertTodoResponse(todo),
	}
	return c.Status(http.StatusOK).JSON(response)
}

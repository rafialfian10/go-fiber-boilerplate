package handlerCategory

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerCategory) CreateCategory(c *fiber.Ctx) error {
	var request dto.CreateCategoryRequest

	err := c.BodyParser(&request)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	category := models.Category{
		Category: request.Category,
	}

	addedCategory, err := h.CategoryRepository.CreateCategory(&category)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	newCategory, err := h.CategoryRepository.GetCategoryByID(addedCategory.ID)
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
		Data:    convertCategoryResponse(newCategory),
	}
	return c.Status(http.StatusCreated).JSON(response)
}

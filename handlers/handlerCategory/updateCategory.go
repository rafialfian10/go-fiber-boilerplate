package handlerCategory

import (
	"go-restapi-boilerplate/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerCategory) UpdateCategory(c *fiber.Ctx) error {
	var request dto.UpdateCategoryRequest

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
			Message: "Invalid category ID",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	category, err := h.CategoryRepository.GetCategoryByID(uint(id))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	if request.Category != "" && request.Category != category.Category {
		category.Category = request.Category
	}

	updatedCategory, err := h.CategoryRepository.UpdateCategory(category)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	category, err = h.CategoryRepository.GetCategoryByID(updatedCategory.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "Category successfully updated",
		Data:    convertCategoryResponse(category),
	}
	return c.Status(http.StatusOK).JSON(response)
}

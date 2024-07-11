package handlerRole

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerRole) CreateRole(c *fiber.Ctx) error {
	var request dto.CreateRoleRequest

	err := c.BodyParser(&request)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	role := models.Role{
		Role: request.Role,
	}

	addedRole, err := h.RoleRepository.CreateRole(&role)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	newRole, err := h.RoleRepository.GetRoleByID(addedRole.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusCreated,
		Message: "OK",
		Data:    convertRoleResponse(newRole),
	}
	return c.Status(http.StatusCreated).JSON(response)
}

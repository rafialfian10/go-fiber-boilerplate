package handlerRole

import (
	"go-restapi-boilerplate/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerRole) DeleteRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Invalid role ID",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	role, err := h.RoleRepository.GetRoleByID(uint(id))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	roleUsed, err := h.RoleRepository.CheckIsRoleUsed(role)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	} else if roleUsed {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Can't delete role, role is used by some user",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	deletedRole, err := h.RoleRepository.DeleteRole(role)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertRoleResponse(deletedRole),
	}
	return c.Status(http.StatusOK).JSON(response)
}

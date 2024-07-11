package handlerUser

import (
	"fmt"
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/pkg/helpers"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *handlerUser) DeleteUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	user, err := h.UserRepository.GetUserByID(id)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if user.Image != "" {
		if !helpers.DeleteFile(user.Image) {
			fmt.Println(err.Error())
		}
	}

	deletedUser, err := h.UserRepository.DeleteUser(user, id)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "User deleted successfully",
		Data:    convertUserResponse(deletedUser),
	}
	return c.Status(http.StatusOK).JSON(response)
}

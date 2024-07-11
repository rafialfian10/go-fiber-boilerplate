package handlerAuth

import (
	"go-restapi-boilerplate/dto"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (h *handlerAuth) CheckAuth(c *fiber.Ctx) error {
	// Get JWT payload
	claims, ok := c.Locals("userData").(jwt.MapClaims)
	if !ok {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User data from jwt payload is not found",
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// Log the claims
	// fmt.Println("Claims received in CheckAuth:", claims)

	// Extract user data from JWT claims
	id, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	user, err := h.UserRepository.GetUserByID(id)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User not found",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertLoginResponse(user, strings.Replace(c.Get("Authorization"), "Bearer ", "", -1)),
	}
	return c.Status(http.StatusOK).JSON(response)
}

package handlerUser

import (
	"go-restapi-boilerplate/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (h *handlerUser) GetProfile(c *fiber.Ctx) error {
	// Retrieve JWT claims from context locals
	userClaims, ok := c.Locals("userData").(jwt.MapClaims)
	if !ok {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Invalid user data in JWT claims",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Extract user ID from claims
	userIdStr, ok := userClaims["id"].(string)
	if !ok {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Invalid user ID format in JWT claims",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Invalid user ID format",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	profile, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "Profile retrieved successfully",
		Data:    convertUserResponse(profile),
	}
	return c.Status(http.StatusOK).JSON(response)
}

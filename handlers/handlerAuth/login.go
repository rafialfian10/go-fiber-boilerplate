package handlerAuth

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/pkg/bcrypt"
	jwtToken "go-restapi-boilerplate/pkg/jwt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (h *handlerAuth) Login(c *fiber.Ctx) error {
	var request dto.LoginRequest

	// Get request data
	err := c.BodyParser(&request)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Check if Email is provided
	if request.Email == "" {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email is required",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Find user data
	user, err := h.UserRepository.GetUserByEmail(request.Email)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User not found",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Check password
	if isPasswordValid := bcrypt.CheckPassword(request.Password, user.Password); !isPasswordValid {
		response := dto.Result{
			Status:  http.StatusUnauthorized,
			Message: "Password invalid",
		}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	// Preparing jwt claims
	myClaims := jwt.MapClaims{}
	myClaims["id"] = user.ID
	myClaims["fullname"] = user.FullName
	myClaims["email"] = user.Email
	myClaims["roleId"] = user.RoleID
	myClaims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24 hours expired

	// Generate token
	token, err := jwtToken.GenerateToken(&myClaims)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// Send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertLoginResponse(user, token),
	}
	return c.Status(http.StatusOK).JSON(response)
}

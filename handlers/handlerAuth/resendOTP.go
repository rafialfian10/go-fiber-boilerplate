package handlerAuth

import (
	"go-restapi-boilerplate/dto"
	"net/http"
	"os"
	"regexp"

	"github.com/asidikrdn/otptimize"
	"github.com/gofiber/fiber/v2"
)

func (h *handlerAuth) ResendOTP(c *fiber.Ctx) error {
	email := c.Params("email")

	// Regular expression pattern for validating email
	emailRegexPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex
	emailRegex := regexp.MustCompile(emailRegexPattern)

	// Validate email format
	if !emailRegex.MatchString(email) {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Invalid email format",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Check if email is registered
	user, err := h.UserRepository.GetUserByEmailOrPhone(email, "")
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email not registered, please register first",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Check if email is already verified
	if user.IsEmailVerified {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email already verified",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Generate and send OTP
	err = otptimize.GenerateAndSendOTP(6, 7, os.Getenv("APP_NAME"), user.FullName, user.Email)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: "Failed to send OTP, please try again later",
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OTP has been sent successfully",
	}
	return c.Status(http.StatusOK).JSON(response)
}

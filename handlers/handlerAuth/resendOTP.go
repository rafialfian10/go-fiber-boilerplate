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

	// Ekspresi reguler untuk validasi email
	regexStr := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile regex
	regex := regexp.MustCompile(regexStr)

	// check email
	if !regex.MatchString(email) {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email invalid",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	user, err := h.UserRepository.GetUserByEmail(email)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email not registered, please register first",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if user.IsEmailVerified {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email already verified",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	otptimize.GenerateAndSendOTP(6, 7, os.Getenv("APP_NAME"), user.FullName, user.Email)

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OTP has been sent successfully",
	}
	return c.Status(http.StatusOK).JSON(response)
}

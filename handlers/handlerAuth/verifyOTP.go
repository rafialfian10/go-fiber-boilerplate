package handlerAuth

import (
	"go-restapi-boilerplate/dto"
	"net/http"

	"github.com/asidikrdn/otptimize"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerAuth) VerifyEmail(c *fiber.Ctx) error {
	var request dto.VerifyEmailRequest
	if err := c.BodyParser(&request); err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	user, err := h.UserRepository.GetUserByEmail(request.Email)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User not found",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	isOtpValid, err := otptimize.ValidateOTP(request.Email, request.OTPToken)
	if !isOtpValid {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "OTP invalid",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	} else if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	user.IsEmailVerified = true

	_, err = h.UserRepository.UpdateUser(user)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "User's email verified",
		Data:    convertRegisterResponse(user),
	}
	return c.Status(http.StatusOK).JSON(response)
}

package handlerAuth

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/pkg/bcrypt"
	"net/http"
	"os"

	"github.com/asidikrdn/otptimize"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *handlerAuth) RegisterUser(c *fiber.Ctx) error {
	var request dto.RegisterRequest

	if err := c.BodyParser(&request); err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Check email and phone
	user, err := h.UserRepository.GetUserByEmailOrPhone(request.Email, request.Phone)
	if err == nil {
		if user.Email == request.Email {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: "Email already registered",
			}
			return c.Status(http.StatusBadRequest).JSON(response)
		}
		if user.Phone == request.Phone {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: "Phone number already registered",
			}
			return c.Status(http.StatusBadRequest).JSON(response)
		}
	}

	newUser := models.User{
		ID:              uuid.New(),
		FullName:        request.FullName,
		Email:           request.Email,
		IsEmailVerified: false,
		Phone:           request.Phone,
		IsPhoneVerified: false,
		RoleID:          3,
	}

	// Hashing password
	newUser.Password, err = bcrypt.HashingPassword(request.Password)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	addedUser, err := h.UserRepository.CreateUser(&newUser)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// Generate and send OTP
	go otptimize.GenerateAndSendOTP(6, 7, os.Getenv("APP_NAME"), newUser.FullName, newUser.Email)

	createdUser, err := h.UserRepository.GetUserByID(addedUser.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusCreated,
		Message: "Register successfully",
		Data:    convertRegisterResponse(createdUser),
	}
	return c.Status(http.StatusCreated).JSON(response)
}

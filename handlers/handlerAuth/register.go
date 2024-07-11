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

	// get request data
	if err := c.BodyParser(&request); err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// check email
	_, err := h.UserRepository.GetUserByEmail(request.Email)
	if err == nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email already registered",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// check phone
	_, err = h.UserRepository.GetUserByEmail(request.Phone)
	if err == nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Phone number already registered",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// create new user
	user := models.Users{
		ID:              uuid.New(),
		FullName:        request.FullName,
		Email:           request.Email,
		IsEmailVerified: false,
		Phone:           request.Phone,
		IsPhoneVerified: false,
		RoleID:          3,
	}

	// hashing password
	user.Password, err = bcrypt.HashingPassword(request.Password)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// get image from context
	// image, ok := c.Locals("image").(string)
	// if ok {
	// 	user.Image = image
	// }

	// save new user data to database
	addedUser, err := h.UserRepository.CreateUser(&user)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// generate and send otp
	go otptimize.GenerateAndSendOTP(6, 7, os.Getenv("APP_NAME"), user.FullName, user.Email)

	// reload data
	newUser, err := h.UserRepository.GetUserByID(addedUser.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// send response
	response := dto.Result{
		Status:  http.StatusCreated,
		Message: "OK",
		Data:    convertRegisterResponse(newUser),
	}
	return c.Status(http.StatusCreated).JSON(response)
}

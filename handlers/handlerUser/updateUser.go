package handlerUser

import (
	"fmt"
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/pkg/helpers"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *handlerUser) UpdateUser(c *fiber.Ctx) error {
	var request dto.UpdateUserRequest

	if err := c.BodyParser(&request); err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	id, err := uuid.Parse(c.Params("id"))
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
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	updateFullname(user, request.FullName)
	updateEmail(user, request.Email)
	updatePhone(user, request.Phone)
	updateGender(user, request.Gender)
	updateAddress(user, request.Address)

	image, ok := c.Locals("image").(string)
	if ok && image != "" {
		if user.Image != "" {
			if !helpers.DeleteFile(user.Image) {
				fmt.Println("Failed to delete image file")
			}
		}
		user.Image = image
	}

	user, err = h.UserRepository.UpdateUser(user)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	user, err = h.UserRepository.GetUserByID(user.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertUserResponse(user),
	}
	return c.Status(http.StatusOK).JSON(response)
}

func (h *handlerUser) UpdateUserByAdmin(c *fiber.Ctx) error {
	var request dto.UpdateUserRequest

	if err := c.BodyParser(&request); err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	id, err := uuid.Parse(c.Params("id"))
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
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	updateFullname(user, request.FullName)
	updateEmail(user, request.Email)
	updatePhone(user, request.Phone)
	updateGender(user, request.Gender)
	updateAddress(user, request.Address)
	updateRole(user, request.RoleID)

	user, err = h.UserRepository.UpdateUserByAdmin(user)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	user, err = h.UserRepository.GetUserByID(user.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertUserResponse(user),
	}
	return c.Status(http.StatusOK).JSON(response)
}

func updateFullname(user *models.User, requestData string) {
	if requestData != "" && requestData != user.FullName {
		user.FullName = requestData
	}
}

func updateEmail(user *models.User, requestData string) {
	if requestData != "" && requestData != user.Email {
		user.IsEmailVerified = false
		user.Email = requestData
	}
}

func updatePhone(user *models.User, requestData string) {
	if requestData != "" && requestData != user.Phone {
		user.IsPhoneVerified = false
		user.Phone = requestData
	}
}

func updateGender(user *models.User, requestData string) {
	if requestData != "" && requestData != user.Gender {
		user.Gender = requestData
	}
}

func updateAddress(user *models.User, requestData string) {
	if requestData != "" && requestData != user.Address {
		user.Address = requestData
	}
}

func updateRole(user *models.User, requestData uint) {
	if requestData != 0 && requestData != user.RoleID {
		user.RoleID = requestData
	}
}

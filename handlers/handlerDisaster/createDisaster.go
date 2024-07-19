package handlerDisaster

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (h *handlerDisaster) CreateDisaster(c *fiber.Ctx) error {
	var request dto.CreateDisasterRequest

	err := c.BodyParser(&request)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	claims, ok := c.Locals("userData").(jwt.MapClaims)
	if !ok {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User data from jwt payload is not found",
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// Extract user data from JWT claims
	userId, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	date, _ := time.Parse("2006-01-02", c.FormValue("date"))

	image, ok := c.Locals("image").(string)
	if !ok {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Image is not provided",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	disaster := models.Disaster{
		UserID:       userId,
		Title:        request.Title,
		Description:  request.Description,
		Location:     request.Location,
		CategoryID:   request.CategoryID,
		Date:         date,
		Donate:       request.Donate,
		DonateTarget: request.DonateTarget,
		Image:        image,
		IsTrending:   request.IsTrending,
	}

	addedDisaster, err := h.DisasterRepository.CreateDisaster(&disaster)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	newDisaster, err := h.DisasterRepository.GetDisasterByID(addedDisaster.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusCreated,
		Message: "Disaster successfully created",
		Data:    convertDisasterResponse(newDisaster),
	}
	return c.Status(http.StatusCreated).JSON(response)
}

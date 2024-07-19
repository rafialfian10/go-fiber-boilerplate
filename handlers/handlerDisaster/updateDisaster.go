package handlerDisaster

import (
	"fmt"
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/pkg/helpers"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerDisaster) UpdateDisaster(c *fiber.Ctx) error {
	var request dto.UpdateDisasterRequest

	if err := c.BodyParser(&request); err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Invalid disaster ID",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	disaster, err := h.DisasterRepository.GetDisasterByID(uint(id))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	if request.Title != "" && request.Title != disaster.Title {
		disaster.Title = request.Title
	}

	if request.Description != "" && request.Description != disaster.Description {
		disaster.Description = request.Description
	}

	if request.Location != "" && request.Location != disaster.Location {
		disaster.Location = request.Location
	}

	if request.CategoryID != 0 && request.CategoryID != disaster.CategoryID {
		disaster.CategoryID = request.CategoryID
	}

	if request.Date != "" {
		date, err := time.Parse("2006-01-02", request.Date)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: "Invalid date format",
			}
			return c.Status(http.StatusBadRequest).JSON(response)
		}
		disaster.Date = date
	}

	if request.Donate != 0 && request.Donate != disaster.Donate {
		disaster.Donate = request.Donate
	}

	if request.DonateTarget != 0 && request.DonateTarget != disaster.DonateTarget {
		disaster.DonateTarget = request.DonateTarget
	}

	image, ok := c.Locals("image").(string)
	if ok && image != "" {
		if disaster.Image != "" {
			if !helpers.DeleteFile(disaster.Image) {
				fmt.Println("Failed to delete image file")
			}
		}
		disaster.Image = image
	}

	disaster.IsTrending = request.IsTrending

	updatedDisaster, err := h.DisasterRepository.UpdateDisaster(disaster)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	disaster, err = h.DisasterRepository.GetDisasterByID(updatedDisaster.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "Disaster successfully updated",
		Data:    convertDisasterResponse(disaster),
	}
	return c.Status(http.StatusOK).JSON(response)
}

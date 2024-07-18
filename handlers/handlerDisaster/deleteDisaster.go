package handlerDisaster

import (
	"fmt"
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerDisaster) DeleteDisaster(c *fiber.Ctx) error {
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

	if disaster.Image != "" {
		if !helpers.DeleteFile(disaster.Image) {
			fmt.Println(err.Error())
		}
	}

	deletedDisaster, err := h.DisasterRepository.DeleteDisaster(disaster)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "Disaster successfully deleted",
		Data:    convertDisasterResponse(deletedDisaster),
	}
	return c.Status(http.StatusOK).JSON(response)
}

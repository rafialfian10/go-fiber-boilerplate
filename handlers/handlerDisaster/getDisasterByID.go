package handlerDisaster

import (
	"go-restapi-boilerplate/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerDisaster) GetDisasterByID(c *fiber.Ctx) error {
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

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertDisasterResponse(disaster),
	}
	return c.Status(http.StatusOK).JSON(response)
}

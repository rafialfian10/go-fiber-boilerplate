package handlerDisaster

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerDisaster) GetDisasters(c *fiber.Ctx) error {
	var (
		disasters     *[]models.Disaster
		err           error
		totalDisaster int64
	)

	// get search query
	searchQuery := c.Query("search")

	// with pagination
	if pageStr := c.Query("page"); pageStr != "" {
		var (
			limit  int
			offset int
		)

		// get page position
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
			return c.Status(http.StatusBadRequest).JSON(response)
		}

		// set limit (if not exist, use default limit -> 5)
		if limitStr := c.Query("limit"); limitStr != "" {
			limit, err = strconv.Atoi(limitStr)
			if err != nil {
				response := dto.Result{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
				}
				return c.Status(http.StatusBadRequest).JSON(response)
			}
		} else {
			limit = 5
		}

		// set offset
		if page == 1 {
			offset = -1
		} else {
			offset = (page * limit) - limit
		}

		disasters, totalDisaster, err = h.DisasterRepository.GetDisasters(limit, offset, searchQuery)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			}
			return c.Status(http.StatusNotFound).JSON(response)
		}

		response := dto.Result{
			Status:      http.StatusOK,
			Message:     "OK",
			TotalData:   totalDisaster,
			TotalPages:  int(math.Ceil(float64(totalDisaster) / float64(limit))),
			CurrentPage: page,
			Data:        convertMultipleDisasterResponse(disasters),
		}
		return c.Status(http.StatusOK).JSON(response)
	}

	// without pagination
	disasters, totalDisaster, err = h.DisasterRepository.GetDisasters(-1, -1, searchQuery)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	response := dto.Result{
		Status:      http.StatusOK,
		Message:     "OK",
		TotalData:   totalDisaster,
		TotalPages:  1,
		CurrentPage: 1,
		Data:        convertMultipleDisasterResponse(disasters),
	}
	return c.Status(http.StatusOK).JSON(response)
}

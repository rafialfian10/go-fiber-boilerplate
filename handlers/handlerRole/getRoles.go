package handlerRole

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerRole) GetRoles(c *fiber.Ctx) error {
	var (
		roles     *[]models.Role
		err       error
		totalRole int64
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

		// get role data from database
		roles, totalRole, err = h.RoleRepository.GetRoles(limit, offset, searchQuery)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			}
			return c.Status(http.StatusNotFound).JSON(response)
		}

		// send response
		response := dto.Result{
			Status:      http.StatusOK,
			Message:     "OK",
			TotalData:   totalRole,
			TotalPages:  int(math.Ceil(float64(totalRole) / float64(limit))),
			CurrentPage: page,
			Data:        convertMultipleRoleResponse(roles),
		}
		return c.Status(http.StatusOK).JSON(response)
	}

	// without pagination
	// get role data from database
	roles, totalRole, err = h.RoleRepository.GetRoles(-1, -1, searchQuery)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	// send response
	response := dto.Result{
		Status:      http.StatusOK,
		Message:     "OK",
		TotalData:   totalRole,
		TotalPages:  1,
		CurrentPage: 1,
		Data:        convertMultipleRoleResponse(roles),
	}
	return c.Status(http.StatusOK).JSON(response)
}

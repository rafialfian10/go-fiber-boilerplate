package handlerUser

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerUser) GetUsers(c *fiber.Ctx) error {
	var (
		users       *[]models.User
		err         error
		totalUser   int64
		filterQuery dto.UserFilter
	)

	// Get filter data
	roleId, _ := strconv.Atoi(c.Query("roleId"))
	filterQuery.RoleID = uint(roleId)

	// Get search query
	searchQuery := c.Query("search")

	// With pagination
	if c.Query("page") != "" {
		var (
			limit  int
			offset int
		)

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
			return c.Status(http.StatusBadRequest).JSON(response)
		}

		// Set limit (if not exist, use default limit -> 5)
		if c.Query("limit") != "" {
			limit, err = strconv.Atoi(c.Query("limit"))
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

		// Set offset
		if page == 1 {
			offset = -1
		} else {
			offset = (page * limit) - limit
		}

		users, totalUser, err = h.UserRepository.GetUsers(limit, offset, filterQuery, searchQuery)
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
			TotalData:   totalUser,
			TotalPages:  int(math.Ceil(float64(float64(totalUser) / float64(limit)))),
			CurrentPage: page,
			Data:        convertMultipleUserResponse(users),
		}
		return c.Status(http.StatusOK).JSON(response)
	} else { // Without pagination
		users, totalUser, err = h.UserRepository.GetUsers(-1, -1, filterQuery, searchQuery)
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
			TotalData:   totalUser,
			TotalPages:  1,
			CurrentPage: 1,
			Data:        convertMultipleUserResponse(users),
		}
		return c.Status(http.StatusOK).JSON(response)
	}
}

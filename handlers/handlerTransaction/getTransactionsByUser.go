package handlerTransaction

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (h *handlerTransaction) GetTransactionsByUser(c *fiber.Ctx) error {
	var (
		transactions     *[]models.Transaction
		err              error
		totalTransaction int64
	)

	// Get authenticated user's ID from JWT claims
	claims, ok := c.Locals("userData").(jwt.MapClaims)
	if !ok {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User data from jwt payload is not found",
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	userID, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// Get search query
	searchQuery := c.Query("search")

	// With pagination
	if pageStr := c.Query("page"); pageStr != "" {
		var (
			limit  int
			offset int
		)

		// Get page position
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
			return c.Status(http.StatusBadRequest).JSON(response)
		}

		// Set limit (if not exist, use default limit -> 5)
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

		// Set offset
		if page == 1 {
			offset = 0
		} else {
			offset = (page * limit) - limit
		}

		// Retrieve transactions by user ID
		transactions, totalTransaction, err = h.TransactionRepository.GetTransactionsByUser(userID, limit, offset, searchQuery)
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
			TotalData:   totalTransaction,
			TotalPages:  int(math.Ceil(float64(totalTransaction) / float64(limit))),
			CurrentPage: page,
			Data:        convertMultipleTransactionResponse(transactions),
		}
		return c.Status(http.StatusOK).JSON(response)
	}

	// Without pagination
	transactions, totalTransaction, err = h.TransactionRepository.GetTransactionsByUser(userID, -1, -1, searchQuery)
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
		TotalData:   totalTransaction,
		TotalPages:  1,
		CurrentPage: 1,
		Data:        convertMultipleTransactionResponse(transactions),
	}
	return c.Status(http.StatusOK).JSON(response)
}

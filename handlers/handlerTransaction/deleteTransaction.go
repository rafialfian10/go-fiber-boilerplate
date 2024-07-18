package handlerTransaction

import (
	"go-restapi-boilerplate/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerTransaction) DeleteTransaction(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Invalid transaction ID",
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	transaction, err := h.TransactionRepository.GetTransactionByID(uint(id))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	deletedTransaction, err := h.TransactionRepository.DeleteTransaction(transaction)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "Transaction successfully deleted",
		Data:    convertTransactionResponse(deletedTransaction),
	}
	return c.Status(http.StatusOK).JSON(response)
}

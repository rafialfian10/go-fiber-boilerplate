package handlerTransaction

import (
	"go-restapi-boilerplate/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerTransaction) UpdateTransaction(c *fiber.Ctx) error {
	var request dto.UpdateTransactionRequest

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

	if request.Status != "" && request.Status != transaction.Status {
		transaction.Status = request.Status
	}

	updatedTransaction, err := h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	transaction, err = h.TransactionRepository.GetTransactionByID(updatedTransaction.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "Transaction successfully updated",
		Data:    convertTransactionResponse(transaction),
	}
	return c.Status(http.StatusOK).JSON(response)
}

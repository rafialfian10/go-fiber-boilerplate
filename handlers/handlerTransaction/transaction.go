package handlerTransaction

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/repositories"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(transactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{transactionRepository}
}

func convertTransactionResponse(transaction *models.Transaction) *dto.TransactionResponse {
	return &dto.TransactionResponse{
		ID:         transaction.ID,
		DisasterID: transaction.DisasterID,
		Disaster: dto.DisasterResponse{
			ID:           transaction.Disaster.ID,
			Title:        transaction.Disaster.Title,
			Description:  transaction.Disaster.Description,
			Location:     transaction.Disaster.Location,
			CategoryID:   transaction.Disaster.CategoryID,
			Category:     transaction.Disaster.Category.Category,
			Date:         transaction.Disaster.Date.Format("2006-01-02"),
			Donate:       transaction.Disaster.Donate,
			DonateTarget: transaction.Disaster.DonateTarget,
			Image:        transaction.Disaster.Image,
			IsTrending:   transaction.Disaster.IsTrending,
			UserID:       transaction.Disaster.UserID,
			User: dto.UserResponse{
				ID:       transaction.Disaster.User.ID,
				FullName: transaction.Disaster.User.FullName,
				Role: dto.RoleResponse{
					ID:   transaction.Disaster.User.Role.ID,
					Role: transaction.Disaster.User.Role.Role,
				},
			},
		},
		Status:          transaction.Status,
		TransactionDate: transaction.TransactionDate.Format("2006-01-02"),
		Token:           transaction.Token,
		UserID:          transaction.UserID,
		User: dto.UserResponse{
			ID:       transaction.User.ID,
			FullName: transaction.User.FullName,
			Email:    transaction.User.Email,
			Phone:    transaction.User.Phone,
			Gender:   transaction.User.Gender,
			Address:  transaction.User.Address,
			Image:    transaction.User.Image,
			Role: dto.RoleResponse{
				ID:   transaction.User.Role.ID,
				Role: transaction.User.Role.Role,
			},
		},
	}
}

func convertMultipleTransactionResponse(transactionDatas *[]models.Transaction) *[]dto.TransactionResponse {
	var transactions []dto.TransactionResponse

	for _, t := range *transactionDatas {
		transactions = append(transactions, *convertTransactionResponse(&t))
	}

	return &transactions
}

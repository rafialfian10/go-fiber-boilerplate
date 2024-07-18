package dto

import "github.com/google/uuid"

type CreateTransactionRequest struct {
	DisasterID      uint   `json:"disaster_id" form:"disaster_id" binding:"required"`
	TransactionDate string `json:"transaction_date" form:"transaction_date" binding:"required"`
	Status          string `json:"status" form:"status" binding:"required"`
}

type UpdateTransactionRequest struct {
	Status string `json:"status" form:"status"`
}

type TransactionResponse struct {
	ID              uint             `json:"id"`
	UserID          uuid.UUID        `json:"user_id"`
	User            UserResponse     `json:"user"`
	DisasterID      uint             `json:"disaster_id"`
	Disaster        DisasterResponse `json:"disaster"`
	Status          string           `json:"status"`
	TransactionDate string           `json:"transaction_date"`
	Token           string           `json:"token"`
}

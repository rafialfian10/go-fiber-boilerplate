package dto

import "github.com/google/uuid"

type CreateDisasterRequest struct {
	Title        string `json:"title" form:"title" binding:"required"`
	Description  string `json:"description" form:"description" binding:"required"`
	Location     string `json:"location" form:"location" binding:"required"`
	CategoryID   uint   `json:"category_id" form:"category_id" binding:"required"`
	Date         string `json:"date" form:"date" binding:"required"`
	Donate       int    `json:"donate" form:"donate"`
	DonateTarget int    `json:"donate_target" form:"donate_target"`
	Image        string `json:"image" form:"image"`
	IsTrending   bool   `json:"is_trending" form:"is_trending"`
}

type UpdateDisasterRequest struct {
	Title        string `json:"title" form:"title"`
	Description  string `json:"description" form:"description"`
	Location     string `json:"location" form:"location"`
	CategoryID   uint   `json:"category_id" form:"category_id"`
	Date         string `json:"date" form:"date"`
	Donate       int    `json:"donate" form:"donate"`
	DonateTarget int    `json:"donate_target" form:"donate_target"`
	Image        string `json:"image" form:"image"`
	IsTrending   bool   `json:"is_trending" form:"is_trending"`
}

type DisasterResponse struct {
	ID           uint         `json:"id"`
	UserID       uuid.UUID    `json:"user_id"`
	User         UserResponse `json:"user"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Location     string       `json:"location"`
	CategoryID   uint         `json:"category_id"`
	Category     string       `json:"category"`
	Date         string       `json:"date"`
	Donate       int          `json:"donate"`
	DonateTarget int          `json:"donate_target"`
	Image        string       `json:"image"`
	IsTrending   bool         `json:"is_trending"`
}

package dto

import (
	"github.com/google/uuid"
)

type UpdateUserRequest struct {
	FullName string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Gender   string `json:"gender" form:"gender"`
	Address  string `json:"address" form:"address"`
	RoleID   uint   `json:"roleId" form:"roleId"`
}

type UserResponse struct {
	ID              uuid.UUID    `json:"id,omitempty"`
	FullName        string       `json:"fullname,omitempty"`
	Email           string       `json:"email,omitempty"`
	IsEmailVerified bool         `json:"isEmailVerified,omitempty"`
	Phone           string       `json:"phone,omitempty"`
	IsPhoneVerified bool         `json:"isPhoneVerified,omitempty"`
	Gender          string       `json:"gender,omitempty"`
	Address         string       `json:"address,omitempty"`
	Image           string       `json:"image,omitempty"`
	Role            RoleResponse `json:"role"`
}

type UserFilter struct {
	RoleID uint
}

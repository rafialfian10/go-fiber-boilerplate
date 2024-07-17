package dto

type CreateCategoryRequest struct {
	Category string `json:"category" form:"category"`
}

type UpdateCategoryRequest struct {
	Category string `json:"category" form:"category"`
}

type CategoryResponse struct {
	ID       uint   `json:"id,omitempty"`
	Category string `json:"category"`
}

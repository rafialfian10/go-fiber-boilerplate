package dto

type CreateTodoRequest struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	CategoryID  uint   `json:"category_id" form:"category_id" binding:"required"`
	Date        string `json:"date" form:"date" binding:"required"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	CategoryID  uint   `json:"category_id" form:"category_id"`
	IsDone      bool   `json:"is_done" form:"is_done"`
	Date        string `json:"date" form:"date"`
}

type TodoResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
	Category    string `json:"category"`
	IsDone      bool   `json:"is_done"`
	Date        string `json:"date"`
}

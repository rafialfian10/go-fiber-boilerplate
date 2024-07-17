package handlerTodo

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/repositories"
)

type handlerTodo struct {
	TodoRepository repositories.TodoRepository
}

func HandlerTodo(todoRepository repositories.TodoRepository) *handlerTodo {
	return &handlerTodo{todoRepository}
}

func convertTodoResponse(todo *models.Todo) *dto.TodoResponse {
	return &dto.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CategoryID:  todo.CategoryID,
		Category:    todo.Category.Category,
		IsDone:      todo.IsDone,
		Date:        todo.Date.Format("2006-01-02"), // Format the date to string
	}
}

func convertMultipleTodoResponse(todoDatas *[]models.Todo) *[]dto.TodoResponse {
	var todos []dto.TodoResponse

	for _, t := range *todoDatas {
		todos = append(todos, *convertTodoResponse(&t))
	}

	return &todos
}

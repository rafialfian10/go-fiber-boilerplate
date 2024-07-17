package repositories

import (
	"fmt"
	"go-restapi-boilerplate/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	GetTodos(limit, offset int, searchQuery string) (*[]models.Todo, int64, error)
	GetTodoByID(todoID uint) (*models.Todo, error)
	CreateTodo(todo *models.Todo) (*models.Todo, error)
	UpdateTodo(todo *models.Todo) (*models.Todo, error)
	DeleteTodo(todo *models.Todo) (*models.Todo, error)
}

func (r *repository) GetTodos(limit, offset int, searchQuery string) (*[]models.Todo, int64, error) {
	var (
		todos     []models.Todo
		totalTodo int64
	)

	trx := r.db.Session(&gorm.Session{})

	if searchQuery != "" {
		trx = trx.Where("todo LIKE ?", fmt.Sprintf("%%%s%%", searchQuery))
	}

	trx.Model(&models.Todo{}).Count(&totalTodo)

	err := trx.Limit(limit).Offset(offset).Preload("Category").Find(&todos).Error

	return &todos, totalTodo, err
}

func (r *repository) GetTodoByID(todoID uint) (*models.Todo, error) {
	var category models.Todo
	err := r.db.Where("id = ?", todoID).Preload("Category").First(&category).Error

	return &category, err
}

func (r *repository) CreateTodo(todo *models.Todo) (*models.Todo, error) {
	err := r.db.Preload("Category").Create(todo).Error

	return todo, err
}

func (r *repository) UpdateTodo(todo *models.Todo) (*models.Todo, error) {
	err := r.db.Model(todo).Preload("Category").Updates(*todo).Error

	return todo, err
}

func (r *repository) DeleteTodo(todo *models.Todo) (*models.Todo, error) {
	err := r.db.Preload("Category").Delete(todo).Error

	return todo, err
}

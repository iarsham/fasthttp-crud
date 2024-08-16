package domain

import (
	"github.com/iarsham/fasthttp-crud/internal/entities"
	"github.com/iarsham/fasthttp-crud/internal/models"
)

type TaskRepository interface {
	Get(id string) (*models.Tasks, error)
	Create(task *entities.TaskRequest) (*models.Tasks, error)
	Update(task *entities.TaskRequest, id string) (*models.Tasks, error)
	Delete(id string) error
}

type TaskService interface {
	GetTask(id string) (*models.Tasks, error)
	CreateTask(task *entities.TaskRequest) (*models.Tasks, error)
	UpdateTask(task *entities.TaskRequest, id string) (*models.Tasks, error)
	DeleteTask(id string) error
}

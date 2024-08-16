package services

import (
	"github.com/iarsham/fasthttp-crud/internal/domain"
	"github.com/iarsham/fasthttp-crud/internal/entities"
	"github.com/iarsham/fasthttp-crud/internal/models"
	"go.uber.org/zap"
)

type taskServiceImpl struct {
	taskRepository domain.TaskRepository
	logger         *zap.Logger
}

func NewTaskService(taskRepository domain.TaskRepository, logger *zap.Logger) domain.TaskService {
	return &taskServiceImpl{
		taskRepository: taskRepository,
		logger:         logger,
	}
}

func (s *taskServiceImpl) GetTask(id string) (*models.Tasks, error) {
	task, err := s.taskRepository.Get(id)
	if err != nil {
		s.logger.Error("Failed to get task", zap.Error(err))
		return nil, err
	}
	return task, nil
}

func (s *taskServiceImpl) CreateTask(task *entities.TaskRequest) (*models.Tasks, error) {
	createdTask, err := s.taskRepository.Create(task)
	if err != nil {
		s.logger.Error("Failed to create task", zap.Error(err))
		return nil, err
	}
	return createdTask, nil
}

func (s *taskServiceImpl) UpdateTask(task *entities.TaskRequest, id string) (*models.Tasks, error) {
	updatedTask, err := s.taskRepository.Update(task, id)
	if err != nil {
		s.logger.Error("Failed to update task", zap.Error(err))
		return nil, err
	}
	return updatedTask, nil
}

func (s *taskServiceImpl) DeleteTask(id string) error {
	if err := s.taskRepository.Delete(id); err != nil {
		s.logger.Error("Failed to delete task", zap.Error(err))
		return err
	}
	return nil
}

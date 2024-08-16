package repository

import (
	"database/sql"
	"github.com/iarsham/fasthttp-crud/internal/domain"
	"github.com/iarsham/fasthttp-crud/internal/entities"
	"github.com/iarsham/fasthttp-crud/internal/models"
)

type taskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) domain.TaskRepository {
	return &taskRepositoryImpl{
		db: db,
	}
}

func (r *taskRepositoryImpl) Get(id string) (*models.Tasks, error) {
	var task models.Tasks
	query := "SELECT * FROM tasks WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.IsDone)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepositoryImpl) Create(task *entities.TaskRequest) (*models.Tasks, error) {
	query := "INSERT INTO tasks (title, is_done) VALUES (?, ?)"
	result, err := r.db.Exec(query, task.Title, task.IsDone)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	task.ID = id
	return (*models.Tasks)(task), nil
}

func (r *taskRepositoryImpl) Update(task *entities.TaskRequest, id string) (*models.Tasks, error) {
	query := "UPDATE tasks SET title = ?, is_done = ? WHERE id = ?"
	_, err := r.db.Exec(query, task.Title, task.IsDone, id)
	if err != nil {
		return nil, err
	}
	return (*models.Tasks)(task), nil
}

func (r *taskRepositoryImpl) Delete(id string) error {
	query := `DELETE FROM tasks WHERE id = ?`
	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}

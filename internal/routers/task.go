package routers

import (
	"database/sql"
	"github.com/fasthttp/router"
	"github.com/iarsham/fasthttp-crud/internal/handlers"
	"github.com/iarsham/fasthttp-crud/internal/repository"
	"github.com/iarsham/fasthttp-crud/internal/services"
	"go.uber.org/zap"
)

func SetupRouter(db *sql.DB, logger *zap.Logger) *router.Router {
	r := router.New()
	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo, logger)
	taskHandler := &handlers.TaskHandler{
		Service: taskService,
	}
	r.GET("/tasks/{id}", taskHandler.GetTaskHandler)
	r.POST("/tasks", taskHandler.CreateTaskHandler)
	r.PUT("/tasks/{id}", taskHandler.UpdateTaskHandler)
	r.DELETE("/tasks/{id}", taskHandler.DeleteTaskHandler)
	return r
}

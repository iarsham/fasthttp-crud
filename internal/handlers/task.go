package handlers

import (
	"database/sql"
	"errors"
	"github.com/iarsham/fasthttp-crud/internal/domain"
	"github.com/iarsham/fasthttp-crud/internal/entities"
	"github.com/iarsham/fasthttp-crud/internal/helpers"
	"github.com/valyala/fasthttp"
)

type TaskHandler struct {
	Service domain.TaskService
}

func (h *TaskHandler) GetTaskHandler(ctx *fasthttp.RequestCtx) {
	taskID := ctx.UserValue("id").(string)
	task, err := h.Service.GetTask(taskID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			helpers.JsonWrite(ctx, fasthttp.StatusNotFound, nil)
		default:
			helpers.JsonWrite(ctx, fasthttp.StatusInternalServerError, task)
		}
		return
	}
	helpers.JsonWrite(ctx, fasthttp.StatusOK, task)
}

func (h *TaskHandler) CreateTaskHandler(ctx *fasthttp.RequestCtx) {
	taskReq := new(entities.TaskRequest)
	if err := helpers.JsonRead(ctx, taskReq); err != nil {
		helpers.JsonWrite(ctx, fasthttp.StatusBadRequest, nil)
		return
	}
	createdTask, err := h.Service.CreateTask(taskReq)
	if err != nil {
		helpers.JsonWrite(ctx, fasthttp.StatusInternalServerError, nil)
		return
	}
	helpers.JsonWrite(ctx, fasthttp.StatusCreated, createdTask)
}

func (h *TaskHandler) UpdateTaskHandler(ctx *fasthttp.RequestCtx) {
	taskReq := new(entities.TaskRequest)
	taskID := ctx.UserValue("id").(string)
	if err := helpers.JsonRead(ctx, taskReq); err != nil {
		helpers.JsonWrite(ctx, fasthttp.StatusBadRequest, nil)
		return
	}
	if _, err := h.Service.GetTask(taskID); errors.Is(err, sql.ErrNoRows) {
		helpers.JsonWrite(ctx, fasthttp.StatusNotFound, nil)
		return
	}
	updatedTask, err := h.Service.UpdateTask(taskReq, taskID)
	if err != nil {
		helpers.JsonWrite(ctx, fasthttp.StatusInternalServerError, nil)
		return
	}
	helpers.JsonWrite(ctx, fasthttp.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTaskHandler(ctx *fasthttp.RequestCtx) {
	taskID := ctx.UserValue("id").(string)
	if _, err := h.Service.GetTask(taskID); errors.Is(err, sql.ErrNoRows) {
		helpers.JsonWrite(ctx, fasthttp.StatusNotFound, nil)
		return
	}
	if err := h.Service.DeleteTask(taskID); err != nil {
		helpers.JsonWrite(ctx, fasthttp.StatusInternalServerError, nil)
		return
	}
	helpers.JsonWrite(ctx, fasthttp.StatusNoContent, nil)
}

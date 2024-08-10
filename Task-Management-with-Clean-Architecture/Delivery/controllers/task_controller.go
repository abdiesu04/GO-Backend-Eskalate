package controllers

import (
	"log"
	"net/http"
	"task_manager/domain"
	"task_manager/usecases"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase usecases.TaskUsecase
}

func NewTaskController(taskUsecase usecases.TaskUsecase) *TaskController {
	return &TaskController{
		TaskUsecase: taskUsecase,
	}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var task domain.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.TaskUsecase.CreateTask(ctx, &task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, task)
}

func (c *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := c.TaskUsecase.GetAllTasks(ctx)
	if err != nil {
		log.Printf("Error fetching tasks: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := c.TaskUsecase.GetTaskByID(ctx, id)
	if err != nil {
		log.Printf("Error fetching task: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if task == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask domain.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.TaskUsecase.UpdateTask(ctx, id, &updatedTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedTask)
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.TaskUsecase.DeleteTask(ctx, id); err != nil {
		log.Printf("Error deleting task: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

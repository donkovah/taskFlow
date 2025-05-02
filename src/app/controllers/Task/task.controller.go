package taskController

import (
	"be/src/app/structs"
	"be/src/domain/models"
	"be/src/domain/service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	service   *service.TaskService
	validator *structs.TaskValidator
}

func NewTaskController(service *service.TaskService) *TaskController {
	return &TaskController{
		service:   service,
		validator: structs.NewTaskValidator(),
	}
}
func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.service.GetTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task"})
	}
	c.JSON(http.StatusOK, task)
}

func (tc TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.service.GetTasks(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch task"})
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate
	validationErr := tc.validator.Validate(structs.MapModelTaskToStructsTask(&task))
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	task.Status = models.StatusTodo

	createdTask, err := tc.service.CreateTask(context.Background(), &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusOK, createdTask)
}

func (tc TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var taskBody *models.Task

	if err := c.ShouldBindJSON(&taskBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.service.GetTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
	}
	task.Title = taskBody.Title
	task.Description = taskBody.Description
	task.Deadline = taskBody.Deadline

	updatedTask, err := tc.service.UpdateTask(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (tc TaskController) StartTask(c *gin.Context) {
	taskId := c.Param("taskId")
	task, err := tc.service.GetTask(context.Background(), taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
		return
	}
	var taskBody *models.Task

	if err := c.ShouldBindJSON(&taskBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Status = models.StatusTodo
	updatedTask, err := tc.service.UpdateTask(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (tc TaskController) BlockTask(c *gin.Context) {
	taskId := c.Param("taskId")
	task, err := tc.service.GetTask(context.Background(), taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
		return
	}

	task.Status = models.StatusBlocked
	task.CompletedAt = time.Now()

	updatedTask, err := tc.service.UpdateTask(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (tc TaskController) CompleteTask(c *gin.Context) {
	taskId := c.Param("taskId")
	task, err := tc.service.GetTask(context.Background(), taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
		return
	}

	task.Status = models.StatusCompleted
	task.CompletedAt = time.Now()

	updatedTask, err := tc.service.UpdateTask(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (ts TaskController) DeleteTask(c *gin.Context) {
	id := c.Param(("id"))
	err := ts.service.DeleteTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
	}
	c.JSON(http.StatusNoContent, nil)
}

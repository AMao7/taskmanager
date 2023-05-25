package handlers

import (
	"net/http"
	"strconv"

	"github.com/AMao7/taskmanager/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Handler struct {
	store entity.TaskStore
}

func NewHandler(store entity.TaskStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) CreateTask(c *gin.Context) {
	var task entity.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	err := validate.Struct(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "fields": err.Error()})
		return
	}

	if err := h.store.Create(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) DeleteTask(c *gin.Context) {
	id := c.Param("id") // Get the ID from the URL path
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}

	err = h.store.Delete(uint(uintID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Task deleted successfully"})
}
func (h *Handler) GetTask(c *gin.Context) {
	id := c.Param("id") // Get the ID from the URL path
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := h.store.GetByID(uint(uintID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating task"})
	}

	task, err := h.store.GetByID(uint(uintID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting task"})
		return
	}

	var newTaskData entity.Task
	if err := c.BindJSON(&newTaskData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}
	// Update the task with new data
	task.UpdateFrom(newTaskData)

	if err := h.store.Update(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Task updated successfully"})

}

// if err := h.store.Update(&task); err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating task"})
// 	return
// }

// c.JSON(http.StatusOK, gin.H{"status": "Task updated successfully"})

func (h *Handler) GetAllTasks(c *gin.Context) {
	task, err := h.store.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting tasks"})
		return
	}

	c.JSON(http.StatusOK, task)
}

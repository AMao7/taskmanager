package handlers

import (
	"net/http"

	"github.com/AMao7/taskmanager/pkg/entity"
	"github.com/gin-gonic/gin"
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

	if err := h.store.Create(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

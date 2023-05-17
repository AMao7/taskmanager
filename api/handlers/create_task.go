package handlers

import (
	"net/http"

	"github.com/AMao7/taskmanager/pkg/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// The function takes a *gin.Context as a parameter. This object allows you to access request and response data and methods, and is passed to every HTTP handler in Gin.

func CreateTask(c *gin.Context) {

	// Inside the function, we first retrieve a database connection from the context:

	db := c.MustGet("db").(*gorm.DB)

	// create a new Task object and attempt to populate it with the JSON body from the incoming request:

	var task entity.Task

	// Then, we create a new Task object and attempt to populate it with the JSON body from the incoming request:

	if err := c.ShouldBindJSON(&task); err != nil {
		// If the JSON body doesn't match the Task struct, ShouldBindJSON will return an error, and we send a 400 Bad Request response to the client.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Payload"})
		return
	}
	if err := db.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create task"})
		return
	}
	c.JSON(http.StatusOK, task)
}

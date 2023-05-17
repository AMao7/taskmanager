package routes

import (
	"github.com/AMao7/taskmanager/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/task", handlers.CreateTask)

	// r.GET("/tasks", handlers.GetAllTasks)

	// r.GET("/task/:id", handlers.GetTask)

	// r.PUT("/task/:id", handlers.UpdateTask)

	// r.DELETE("/task/:id", handlers.DeleteTask)

	return r

}

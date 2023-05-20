package main

import (
	"github.com/AMao7/taskmanager/api/handlers"
	"github.com/AMao7/taskmanager/pkg/entity"
	"github.com/AMao7/taskmanager/pkg/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	db := repository.ConnectDatabase()

	store := entity.NewGormTaskStore(db)

	x := handlers.NewHandler(store)

	router := gin.Default()

	router.POST("/tasks", x.CreateTask)
	// router.GET("/tasks", handler.GetAllTasks)
	// router.GET("/tasks/:id", handler.GetTask)
	// router.PUT("/tasks/:id", handler.UpdateTask)
	// router.DELETE("/tasks/:id", handler.DeleteTask)

	// r.POST("/task", handlers.TaskHandler.CreateTask)
	// r.GET("/getalltask", handlers.h.)

	router.Run(":8080")
}

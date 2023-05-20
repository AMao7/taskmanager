package main

import (
	"github.com/AMao7/taskmanager/api/handlers"
	"github.com/AMao7/taskmanager/pkg/entity"
	"github.com/AMao7/taskmanager/pkg/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	// returns a *gorm.DB
	db := repository.ConnectDatabase()

	// returns a gormTaskStore of *gorm.DB
	store := entity.NewGormTaskStore(db)

	// returns a Handler
	x := handlers.NewHandler(store)

	router := gin.Default()

	router.POST("/tasks", x.CreateTask)
	// router.GET("/tasks", handler.GetAllTasks)
	// router.GET("/tasks/:id", handler.GetTask)
	// router.PUT("/tasks/:id", handler.UpdateTask)
	router.DELETE("/tasks/:id", x.DeleteTask)

	// r.POST("/task", handlers.TaskHandler.CreateTask)
	// r.GET("/getalltask", handlers.h.)

	router.Run(":8080")
}

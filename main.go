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
	router.GET("/tasks", x.GetAllTasks)
	router.GET("/tasks/:id", x.GetTask)
	router.PUT("/tasks/:id", x.UpdateTask)
	router.DELETE("/tasks/:id", x.DeleteTask)

	router.Run(":8080")
}

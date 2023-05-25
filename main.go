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

	SetupRouter(x)

}

func SetupRouter(handler *handlers.Handler) *gin.Engine {
	apitest := gin.Default()
	// api.Use(gin.Logger())

	apitest.Use(handlers.LoggingMiddleware, handlers.ErrorHandlingMiddleware) // This applies middleware to all routes
	api := apitest.Group("/api")

	api.POST("/task", handler.CreateTask)
	api.DELETE("/task/:id", handler.DeleteTask)
	api.GET("/task/:id", handler.GetTask)
	api.PUT("/task/:id", handler.UpdateTask)
	api.GET("/task", handler.GetAllTasks)

	apitest.Run(":8080")

	return apitest
}

package main

import (
	"github.com/AMao7/taskmanager/api/handlers"
	"github.com/AMao7/taskmanager/pkg/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	db := repository.ConnectDatabase()
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.POST("/task", handlers.CreateTask)

	r.Run(":8080")
}

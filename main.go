package main

import (
	"github.com/AMao7/taskmanager/api/routes"
	"github.com/AMao7/taskmanager/pkg/repository"
)

func main() {
	repository.ConnectDatabase()

	r := routes.SetupRoutes()
	r.Run()
}

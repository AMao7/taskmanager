package repository

import (
	"fmt"
	"os"

	"github.com/AMao7/taskmanager/pkg/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	username, password, host, port, database := os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", username, password, host, port, database)

	DB, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("Failed to Connect to database")
	}
	DB.AutoMigrate(&entity.Task{})

	psqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to get PSQL DB!")
	}

	err = psqlDB.Ping()
	if err != nil {
		panic("Failed to ping database!")
	}
	return DB
}

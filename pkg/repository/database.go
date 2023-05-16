package repository

import (
	"github.com/AMao7/taskmanager/pkg/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

username, password, host, port, database := os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "user=username password=password host=host dbname=taskmanager port=port sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to database")
	}
	db.AutoMigrate(&entity.User{}, &entity.Task{})

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get SQL DB!")
	}

	err = sqlDB.Ping()
	if err != nil {
		panic("Failed to ping database!")
	}

}

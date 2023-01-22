package dbconfig

import (
	"fmt"
	"log"

	"github.com/fxmbx/go_practice_pr/config"
	"github.com/fxmbx/go_practice_pr/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection(config config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBPort, config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error creating db connectoin: %v", err)
	}
	db.AutoMigrate(&models.User{}, &models.Post{})
	println("Db Connected ⏲️")
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	sqldb, err := db.DB()
	if err != nil {
		log.Fatalf("failed to close connection:  %v", err)
	}
	sqldb.Close()
}

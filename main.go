package main

import (
	"github.com/fxmbx/go_practice_pr/config"
	"github.com/fxmbx/go_practice_pr/dbconfig"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	conf config.Config = config.LoadConfig(".")
	db   *gorm.DB      = dbconfig.SetUpDatabaseConnection(conf)
)

func main() {
	// defer dbconfig.CloseDatabaseConnection(db)
	server := gin.Default()
	server.Run()
}

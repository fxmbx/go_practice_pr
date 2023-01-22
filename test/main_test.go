package test

import (
	"os"
	"testing"

	"github.com/fxmbx/go_practice_pr/config"
	"github.com/fxmbx/go_practice_pr/dbconfig"
	"github.com/fxmbx/go_practice_pr/repository"
	"gorm.io/gorm"
)

var (
	userRepo repository.UserRepository
	postRepo repository.PostRepository
	conf              = config.LoadConfig("../")
	db       *gorm.DB = dbconfig.SetUpDatabaseConnection(conf)
)

func TestMain(m *testing.M) {

	defer dbconfig.CloseDatabaseConnection(db)

	userRepo = repository.NewUserRepository(db)
	postRepo = repository.NewPostRepository(db)
	os.Exit(m.Run())
}

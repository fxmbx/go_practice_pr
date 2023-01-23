package test

import (
	"os"
	"testing"

	"github.com/fxmbx/go_practice_pr/config"
	"github.com/fxmbx/go_practice_pr/dbconfig"
	"github.com/fxmbx/go_practice_pr/repository"
	userservice "github.com/fxmbx/go_practice_pr/services/user_service"
	"gorm.io/gorm"
)

var (
	userRepo repository.UserRepository
	postRepo repository.PostRepository
	UserServ userservice.UserService
	conf              = config.LoadConfig("../")
	db       *gorm.DB = dbconfig.SetUpDatabaseConnection(conf)
)

func TestMain(m *testing.M) {

	defer dbconfig.CloseDatabaseConnection(db)

	userRepo = repository.NewUserRepository(db)
	postRepo = repository.NewPostRepository(db)
	UserServ = userservice.NewuserService(userRepo)
	os.Exit(m.Run())
}

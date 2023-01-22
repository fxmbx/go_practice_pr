package repository

import (
	"github.com/fxmbx/go_practice_pr/models"
	"github.com/fxmbx/go_practice_pr/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user models.User) (models.User, error)
	UpdatetUser(user models.User) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	FindUserByUserID(userID string) (models.User, error)
	DeleteByUserID(userID string) error
	ListAllUsers(pagination utils.Pagination) (*utils.Pagination, error)
	ListAllUsersIncludeDeleted(pagination utils.Pagination) (*utils.Pagination, error)
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(dbConnection *gorm.DB) UserRepository {
	return &userRepository{
		db: dbConnection,
	}
}
func (userRepo *userRepository) InsertUser(user models.User) (models.User, error) {
	user.Password = utils.HashedPassword(user.Password)
	userRepo.db.Save(&user)
	return user, nil
}

func (userRepo *userRepository) UpdatetUser(user models.User) (models.User, error) {
	if len(user.Password) > 1 {
		user.Password = utils.HashedPassword(user.Password)
	} else {
		var tempUser models.User
		userRepo.db.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	userRepo.db.Save(&user)
	return user, nil
}

func (userRepo *userRepository) FindUserByEmail(email string) (models.User, error) {
	var user models.User

	res := userRepo.db.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (userRepo *userRepository) FindUserByUserID(userID string) (models.User, error) {
	var user models.User

	res := userRepo.db.Where("id = ?", userID).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (userRepo *userRepository) DeleteByUserID(userID string) error {
	userRepo.db.Where("id = ?", userID).Delete(&models.User{})
	return nil
}

func (userRepo *userRepository) ListAllUsers(pagination utils.Pagination) (*utils.Pagination, error) {
	var users []*models.User
	userRepo.db.Scopes(Paginate(users, &pagination, userRepo.db)).Find(&users)
	pagination.Data = users
	return &pagination, nil
}

func (userRepo *userRepository) ListAllUsersIncludeDeleted(pagination utils.Pagination) (*utils.Pagination, error) {

	var (
		users []*models.User
	)
	userRepo.db.Scopes(PaginateIncludeDeleted(users, &pagination, userRepo.db)).Find(&users)
	pagination.Data = users

	return &pagination, nil
}

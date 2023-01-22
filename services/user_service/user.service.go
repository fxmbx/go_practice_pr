package userservice

import (
	"github.com/fxmbx/go_practice_pr/dto"
	"github.com/fxmbx/go_practice_pr/dto/userdto"
	"github.com/fxmbx/go_practice_pr/models"
	"github.com/fxmbx/go_practice_pr/repository"
	"github.com/fxmbx/go_practice_pr/utils"
)

type UserService interface {
	CreateUser(registerRequest userdto.RegisterRequest) (*userdto.UserResponse, error)
	UpdateUser(updateUserRequest userdto.UpdateUserRequest) (*userdto.UserResponse, error)
	FindUserByEmail(email string) (*userdto.UserResponse, error)
	FindUserByID(userID string) (*userdto.UserResponse, error)
	RemoveUser(userID string) error
	GetAllUsers(query dto.RequestQuery) (*utils.Pagination, error)
	GetAllUsersIncludingDeletedUsers(query dto.RequestQuery) (*utils.Pagination, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewuserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (uService *userService) CreateUser(registerRequest userdto.RegisterRequest) (*userdto.UserResponse, error) {
	user := models.User{
		UserName:  registerRequest.UserName,
		Email:     registerRequest.FirstName,
		FitstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		Password:  registerRequest.Password,
	}

	createdUser, err := uService.userRepo.InsertUser(user)
	if err != nil {
		return nil, err
	}
	userres := userdto.NewUserResponse(createdUser)
	return &userres, nil

}
func (uService *userService) UpdateUser(updateUserRequest userdto.UpdateUserRequest) (*userdto.UserResponse, error) {
	user := models.User{
		ID:        updateUserRequest.ID,
		UserName:  updateUserRequest.UserName,
		Email:     updateUserRequest.FirstName,
		FitstName: updateUserRequest.FirstName,
		LastName:  updateUserRequest.LastName,
	}

	updatedUser, err := uService.userRepo.UpdatetUser(user)
	if err != nil {
		return nil, err
	}
	userres := userdto.NewUserResponse(updatedUser)
	return &userres, nil

}
func (uService *userService) FindUserByEmail(email string) (*userdto.UserResponse, error) {

	user, err := uService.userRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	userres := userdto.NewUserResponse(user)
	return &userres, nil

}
func (uService *userService) FindUserByID(userID string) (*userdto.UserResponse, error) {
	user, err := uService.userRepo.FindUserByUserID(userID)
	if err != nil {
		return nil, err
	}
	userres := userdto.NewUserResponse(user)
	return &userres, nil
}

func (uService *userService) RemoveUser(userID string) error {
	if err := uService.userRepo.DeleteByUserID(userID); err != nil {
		return err
	}
	return nil
}
func (uService *userService) GetAllUsers(query dto.RequestQuery) (*utils.Pagination, error) {
	paginationReq := utils.Pagination{
		Page:  query.PageID,
		Limit: query.PageSize,
		Sort:  query.SortBy,
	}
	paginatedRes, err := uService.userRepo.ListAllUsers(paginationReq)
	if err != nil {
		return nil, err
	}
	return paginatedRes, nil
}
func (uService *userService) GetAllUsersIncludingDeletedUsers(query dto.RequestQuery) (*utils.Pagination, error) {
	paginationReq := utils.Pagination{
		Page:  query.PageID,
		Limit: query.PageSize,
		Sort:  query.SortBy,
	}
	paginatedRes, err := uService.userRepo.ListAllUsersIncludeDeleted(paginationReq)
	if err != nil {
		return nil, err
	}
	return paginatedRes, nil
}

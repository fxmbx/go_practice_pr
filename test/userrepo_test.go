package test

import (
	"fmt"
	"testing"

	"github.com/fxmbx/go_practice_pr/models"
	"github.com/fxmbx/go_practice_pr/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) models.User {

	user := models.User{
		FitstName: utils.RandomString(6),
		LastName:  utils.RandomString(6),
		Email:     utils.RandomString(6),
		UserName:  utils.RandomString(6),
		Password:  utils.RandomString(6),
	}
	insertedUser, err := userRepo.InsertUser(user)
	require.NoError(t, err)
	require.Equal(t, insertedUser.Email, user.Email)
	return insertedUser
}
func TestInsertUser(t *testing.T) {
	user := CreateRandomUser(t)
	insertedUser, err := userRepo.InsertUser(user)
	require.NoError(t, err)
	require.Equal(t, insertedUser.Email, user.Email)
}

func TestFindUserByEmail(t *testing.T) {
	user := CreateRandomUser(t)
	fetechedUser, err := userRepo.FindUserByEmail(user.Email)
	require.NoError(t, err)
	require.Equal(t, fetechedUser.Email, user.Email)
	require.Equal(t, user, fetechedUser)

}

func TestUpdatetUser(t *testing.T) {
	user := CreateRandomUser(t)
	user.Password = ""
	user.FitstName = "Funbi"

	user, err := userRepo.UpdatetUser(user)
	require.NoError(t, err)
	require.Equal(t, user.FitstName, "Funbi")
	require.NotEmpty(t, user.Password)
}

func TestDeleteByUserID(t *testing.T) {
	user := CreateRandomUser(t)
	err := userRepo.DeleteByUserID(fmt.Sprintf("%d", user.ID))
	require.NoError(t, err)
	fetechedUser, err := userRepo.FindUserByEmail(user.Email)
	require.Error(t, err)
	require.Equal(t, err.Error(), "record not found")
	require.NotNil(t, fetechedUser.Deleted)

}

func TestListAllUsers(t *testing.T) {
	p := utils.Pagination{
		Limit: 10,
		Page:  1,
	}
	paginatedRes, err := userRepo.ListAllUsers(p)
	require.NoError(t, err)
	require.NotEmpty(t, paginatedRes)
	println(paginatedRes.Page)
	println(paginatedRes.TotalRows)
	println(paginatedRes.TotalPages)
	require.Equal(t, paginatedRes.Page, int32(1))

}
func TestListAllUsersIncludeDeleted(t *testing.T) {
	p := utils.Pagination{
		Limit: 10,
		Page:  1,
	}
	paginatedRes, err := userRepo.ListAllUsersIncludeDeleted(p)
	require.NoError(t, err)
	require.NotEmpty(t, paginatedRes)

	println(paginatedRes.Page)
	println(paginatedRes.TotalRows)
	println(paginatedRes.TotalPages)

	require.Equal(t, paginatedRes.Page, int32(1))

}

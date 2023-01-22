package test

import (
	"fmt"
	"testing"

	"github.com/fxmbx/go_practice_pr/models"
	"github.com/fxmbx/go_practice_pr/utils"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) models.Post {
	user := CreateRandomUser(t)
	post := models.Post{
		Title:       utils.RandomString(6),
		Description: utils.RandomString(16),
		UserID:      user.ID,
		User:        user,
	}
	insertedPost, err := postRepo.InserPost(post)
	require.NoError(t, err)
	require.Equal(t, insertedPost.Title, post.Title)
	require.Equal(t, insertedPost.UserID, user.ID)
	return insertedPost
}

func TestInserPost(t *testing.T) {
	user := CreateRandomUser(t)

	post := models.Post{
		Title:       utils.RandomString(6),
		Description: utils.RandomString(16),
		UserID:      user.ID,
		User:        user,
	}
	insertedPost, err := postRepo.InserPost(post)
	require.NoError(t, err)
	require.Equal(t, insertedPost.Title, post.Title)
	require.Equal(t, insertedPost.UserID, user.ID)
}

func TestFindPostByPostID(t *testing.T) {
	post := createRandomPost(t)
	fetechedPost, err := postRepo.FindPostByPostID(fmt.Sprintf("%d", post.ID))
	require.NoError(t, err)
	require.Equal(t, fetechedPost, post)
}
func TestUpdatePost(t *testing.T) {
	post := createRandomPost(t)
	post.Title = "Post one"
	post.Description = "Post one Descriptoin"
	updatedPost, err := postRepo.UpdatePost(post)

	require.NoError(t, err)
	require.Equal(t, updatedPost.Title, post.Title)
	require.Equal(t, updatedPost.Title, post.Title)

}

func TestListPost(t *testing.T) {
	p := utils.Pagination{
		Page:  1,
		Limit: 10,
	}
	paginatedRes, err := postRepo.ListPost(p)
	require.NoError(t, err)

	println(paginatedRes.Page)
	println(paginatedRes.TotalPages)
	println(paginatedRes.TotalRows)
	println(paginatedRes.Data)

	require.True(t, paginatedRes.TotalPages < 10)
}

func TestDeletePost(t *testing.T) {
	post := createRandomPost(t)

	err := postRepo.DeletePost(fmt.Sprintf("%d", post.ID))
	require.NoError(t, err)

	fetchedPost, err := postRepo.FindPostByPostID(fmt.Sprintf("%d", post.ID))

	require.Error(t, err)
	require.Equal(t, err.Error(), "record not found")
	require.NotNil(t, fetchedPost.Deleted)

}

func TestListPostIncludeDeleted(t *testing.T) {
	p := utils.Pagination{
		Limit: 10,
		Page:  1,
	}
	paginatedRes, err := postRepo.ListPostIncludeDeleted(p)
	require.NoError(t, err)
	require.NotEmpty(t, paginatedRes)

	println(paginatedRes.Page)
	println(paginatedRes.TotalRows)
	println(paginatedRes.TotalPages)

	require.Equal(t, paginatedRes.Page, int32(1))
}

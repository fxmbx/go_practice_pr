package repository

import (
	"github.com/fxmbx/go_practice_pr/models"
	"github.com/fxmbx/go_practice_pr/utils"
	"gorm.io/gorm"
)

type PostRepository interface {
	InserPost(post models.Post) (models.Post, error)
	UpdatePost(post models.Post) (models.Post, error)
	FindPostByTitle(title string) (models.Post, error)
	FindPostByPostID(postId string) (models.Post, error)
	ListPost(pagination utils.Pagination) (*utils.Pagination, error)
	DeletePost(postId string) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(dbConnectoin *gorm.DB) PostRepository {
	return &postRepository{
		db: dbConnectoin,
	}
}

func (postRepo *postRepository) InserPost(post models.Post) (models.Post, error) {
	postRepo.db.Save(&post)
	postRepo.db.Preload("User").Find(&post)
	return post, nil
}
func (postRepo *postRepository) UpdatePost(post models.Post) (models.Post, error) {
	postRepo.db.Save(&post)
	postRepo.db.Preload("User").Find(&post)
	return post, nil
}
func (postRepo *postRepository) FindPostByTitle(title string) (models.Post, error) {
	var post models.Post
	res := postRepo.db.Preload("User").Where("title = ?", title).Take(&post)
	if res.Error != nil {
		return post, res.Error
	}
	return post, nil
}
func (postRepo *postRepository) FindPostByPostID(postId string) (models.Post, error) {
	var post models.Post
	res := postRepo.db.Preload("User").Where("id = ?", postId).Take(&post)
	if res.Error != nil {
		return post, res.Error
	}
	return post, nil
}
func (postRepo *postRepository) ListPost(pagination utils.Pagination) (*utils.Pagination, error) {
	var posts []*models.Post
	postRepo.db.Scopes(Paginate(posts, &pagination, postRepo.db)).Unscoped().Find(&posts)
	pagination.Data = posts
	return &pagination, nil
}
func (postRepo *postRepository) DeletePost(postId string) error {
	var post models.Post
	res := postRepo.db.Preload("User").Where("title = ?", postId).Take(&post)
	if res.Error != nil {
		return res.Error
	}
	postRepo.db.Delete(&post)
	return nil
}

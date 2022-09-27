package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
	"boilerplate-api/utils"

	"gorm.io/gorm"
)

type PostsService struct {
	repository repository.PostsRepository
}

func NewPostsService(repository repository.PostsRepository) PostsService {
	return PostsService{
		repository: repository,
	}
}

// WithTrx -> enables repository with transaction
func (c PostsService) WithTrx(trxHandle *gorm.DB) PostsService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

// CreatePosts -> call to create the Post
func (c PostsService) CreatePosts(posts models.Post) error {
	return c.repository.CreatePosts(posts)
}

func (c PostsService) UpdatePosts(posts models.Post) error {
	return c.repository.UpdatePosts(posts)
}

func (c PostsService) GetOnePost(postId int64, userId string) (Posts models.UserPost, err error) {
	return c.repository.GetOnePost(postId, userId)
}

func (c PostsService) GetPost(postId int64) (Posts models.Post, err error) {
	return c.repository.GetPost(postId)
}

func (c PostsService) DeletePosts(ID int64) error {
	return c.repository.DeletePosts(ID)
}

// GetAllPosts -> call to get all the Post
func (c PostsService) GetAllPosts(pagination utils.Pagination) ([]models.Post, int64, error) {
	return c.repository.GetAllPosts(pagination)
}

// CreatorPosts -> call to get all creator posts
func (c PostsService) CreatorPosts(cursorPagination utils.CursorPagination, userId string) ([]models.Post, error) {
	return c.repository.CreatorPosts(cursorPagination, userId)
}

// Get User Feeds -> call to get user feeds
func (c PostsService) GetUserFeeds(cursorPagination utils.CursorPagination, userId string) ([]models.FeedPost, error) {
	return c.repository.GetUserFeed(cursorPagination, userId)
}
func (c PostsService) UploadFile(fileName string) {
	c.repository.UploadFile(fileName)
}

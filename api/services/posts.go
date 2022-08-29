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

func (c PostsService) DeletePosts(ID int64) error {
	return c.repository.DeletePosts(ID)
}

// GetAllPosts -> call to get all the Post
func (c PostsService) GetAllPosts(pagination utils.Pagination) ([]models.Post, int64, error) {
	return c.repository.GetAllPosts(pagination)
}

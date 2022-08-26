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

// CreatePosts -> call to create the Posts
func (c PostsService) CreatePosts(Posts models.Posts) error {
	err := c.repository.CreatePosts(Posts)
	return err
}

func (c PostsService) UpdatePosts(Posts models.Posts) error {
	return c.repository.UpdatePosts(Posts)
}

func (c PostsService) DeletePosts(ID int64) error {
	return c.repository.DeletePosts(ID)
}

// GetAllPost -> call to get all the Post
func (c PostsService) GetAllPosts(pagination utils.Pagination) ([]models.Posts, int64, error) {
	return c.repository.GetAllPosts(pagination)
}

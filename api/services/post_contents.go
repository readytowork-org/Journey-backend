package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
	"boilerplate-api/utils"

	"gorm.io/gorm"
)

// PostContentsService -> struct
type PostContentsService struct {
	repository repository.PostContentsRepository
}

// NewPostContentsService -> creates a new PostContentsservice
func NewPostContentsService(repository repository.PostContentsRepository) PostContentsService {
	return PostContentsService{
		repository: repository,
	}
}

// WithTrx -> enables repository with transaction
func (c PostContentsService) WithTrx(trxHandle *gorm.DB) PostContentsService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

// CreatePostContents -> call to create the PostContents
func (c PostContentsService) CreatePostContents(PostContents []models.PostContents) error {
	err := c.repository.CreatePostContents(PostContents)
	return err
}

// GetAllPostContents -> call to get all the PostContents
func (c PostContentsService) GetAllPostContentss(pagination utils.Pagination) ([]models.PostContents, int64, error) {
	return c.repository.GetAllPostContentss(pagination)
}

func (c PostContentsService) UpdatePostContents(PostContents models.PostContents) error {
	return c.repository.UpdatePostContents(PostContents)
}

func (c PostContentsService) DeletePostContents(ID int64) error {
	return c.repository.DeletePostContents(ID)
}

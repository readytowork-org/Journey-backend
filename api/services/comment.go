package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
	"boilerplate-api/utils"

	"gorm.io/gorm"
)

// CommentService -> struct
type CommentService struct {
	repository repository.CommentRepository
}

// NewCommentService -> creates a new Commentservice
func NewCommentService(repository repository.CommentRepository) CommentService {
	return CommentService{
		repository: repository,
	}
}

// WithTrx -> enables repository with transaction
func (c CommentService) WithTrx(trxHandle *gorm.DB) CommentService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

// CreateComment -> call to create the Comment
func (c CommentService) CreateComment(comment models.Comment) error {
	err := c.repository.CreateComment(comment)
	return err
}

// GetAllComment -> call to get all the Comment
func (c CommentService) GetAllComments(pagination utils.Pagination) ([]models.Comment, int64, error) {
	return c.repository.GetAllComments(pagination)
}

func (c CommentService) UpdateComment(comment models.Comment) error {
	return c.repository.UpdateComment(comment)
}

func (c CommentService) DeleteComment(ID int64) error {
	return c.repository.DeleteComment(ID)
}

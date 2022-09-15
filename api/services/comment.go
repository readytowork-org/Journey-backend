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

func (c CommentService) UpdateComment(comment models.Comment) error {
	return c.repository.UpdateComment(comment)
}

func (c CommentService) DeleteComment(comment models.Comment) error {
	return c.repository.DeleteComment(comment)
}

func (c CommentService) CreateCommentLike(commentLike models.CommentLikes) error {
	return c.repository.CreateCommentLike(commentLike)
}

func (c CommentService) DeleteCommentLike(commentLike models.CommentLikes) error {
	return c.repository.DeleteCommentLike(commentLike)
}

func (c CommentService) GetUserPostComment(pagination utils.CursorPagination, postId int64) ([]models.Comment, int64, error) {
	return c.repository.GetUserPostComment(pagination, postId)
}

func (c CommentService) GetOneUserComment(id int64, userId string) (comment models.UserComment, err error) {
	return c.repository.GetOneUserComment(id, userId)
}

func (c CommentService) GetOneComment(id int64) (comment models.Comment, err error) {
	return c.repository.GetOneComment(id)
}

func (c CommentService) GetUserCommentLike(likes models.CommentLikes) (commentLike models.UserCommentLike, err error) {
	return c.repository.GetUserCommentLike(likes)
}

package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
	"time"

	"gorm.io/gorm"
)

// CommentRepository -> database structure
type CommentRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewCommentRepository -> creates a new Comment repository
func NewCommentRepository(db infrastructure.Database, logger infrastructure.Logger) CommentRepository {
	return CommentRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx enables repository with transaction
func (c CommentRepository) WithTrx(trxHandle *gorm.DB) CommentRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

// Save -> Comment
func (c CommentRepository) CreateComment(comment models.Comment) error {
	return c.db.DB.Create(&comment).Error
}

// Update -> Comment
func (c CommentRepository) UpdateComment(comment models.Comment) error {
	return c.db.DB.Model(&models.Comment{}).
		Where("id = ?", comment.ID).
		Updates(map[string]interface{}{
			"comment": comment.Comment,
		}).Error
}

// Delete -> Comment
func (c CommentRepository) DeleteComment(comment models.Comment) error {
	return c.db.DB.Where("id = ?", comment.ID).Where("user_id = ?", comment.UserId).
		Delete(&models.Comment{}).Error
}

func (c CommentRepository) GetUserPostComment(pagination utils.CursorPagination, postId int64) ([]models.Comment, int64, error) {
	var comment []models.Comment
	parsedCursor, _ := time.Parse(time.RFC3339, pagination.Cursor)

	var totalRows int64 = 0
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Order("created_at desc").Model(&models.Comment{}).
		Joins("left join posts on posts.id = comments.post_id").
		Where("posts.id = ? ", postId).
		Where("comments.created_at < ? ", parsedCursor)

	return comment, totalRows, queryBuilder.Find(&comment).Count(&totalRows).Error
}

func (c CommentRepository) CreateCommentLike(commentLikes models.CommentLikes) error {
	return c.db.DB.Create(&commentLikes).Error
}

func (c CommentRepository) GetOneComment(id int64) (comment models.Comment, err error) {
	return comment, c.db.DB.Model(&models.Comment{}).Where("id = ?", id).First(&comment).Error
}

func (c CommentRepository) DeleteCommentLike(commentLikes models.CommentLikes) error {
	return c.db.DB.Delete(&commentLikes).Error
}

func (c CommentRepository) GetOneUserComment(id int64, userId string) (comment models.UserComment, err error) {
	return comment, c.db.DB.
		Model(&models.Comment{}).
		Select(`comments.*,(SELECT COUNT(comment_id)
	FROM comment_likes JOIN comments p ON p.id = comment_likes.comment_id) like_count,
   IF((SELECT c.user_id FROM comment_likes c WHERE user_id = ?) = ?, TRUE, FALSE) has_liked`, userId, userId).
		Where("id = ? ", id).Find(&comment).Error
}

func (c CommentRepository) GetUserCommentLike(likes models.CommentLikes) (commentLike models.UserCommentLike, err error) {
	err = c.db.DB.Select(`comment_id,
	(SELECT COUNT(comment_id)
	 FROM comment_likes
	 JOIN comments p ON p.id = comment_likes.comment_id) like_count,
	IF((SELECT c.user_id FROM comment_likes c WHERE user_id = ?) = ?, TRUE, FALSE) 
	has_liked`, likes.UserId, likes.UserId).Model(&models.CommentLikes{}).Where("comment_id = ?", likes.CommentId).Error
	return commentLike, err
}

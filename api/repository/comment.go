package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"

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
func (c CommentRepository) DeleteComment(ID int64) error {
	println("this is comment id")
	println(ID)
	return c.db.DB.Where("id = ?", ID).
		Delete(&models.Comment{}).Error
}

// GetAllComment -> Get All Comments
func (c CommentRepository) GetAllComments(pagination utils.Pagination) ([]models.Comment, int64, error) {
	var comments []models.Comment
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.Comment{})

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`Comments`.`name` LIKE ?", searchQuery))
	}

	err := queryBuilder.
		Find(&comments).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return comments, totalRows, err
}

func (c CommentRepository) CreateCommentLike(commentLikes models.CommentLikes) error {
	return c.db.DB.Create(&commentLikes).Error
}

func (c CommentRepository) DeleteCommentLike(commentLikes models.CommentLikes) error {
	return c.db.DB.Delete(&commentLikes).Error
}

func (c CommentRepository) GetOneComment(id int64, userId int64) (comment models.UserComment, err error) {
	return comment, c.db.DB.Model(&models.Comment{}).Select(`comments.*,(SELECT COUNT(comment_id)
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

package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"

	"gorm.io/gorm"
)

type LikesRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewLikesRepository(db infrastructure.Database, logger infrastructure.Logger) LikesRepository {
	return LikesRepository{
		db:     db,
		logger: logger,
	}
}

func (c LikesRepository) WithTrx(trxHandle *gorm.DB) LikesRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context.")
		return c
	}
	c.db.DB = trxHandle
	return c
}

func (c LikesRepository) CreateLikes(likes models.PostLike) error {
	return c.db.DB.Create(&likes).Error

}

func (c LikesRepository) GetUserPostLikes(likes models.PostLike) (postLike models.UserPostLike, err error) {
	err = c.db.DB.Select(`post_id,
	(SELECT COUNT(post_id)
	 FROM post_likes
			  JOIN posts p ON p.id = post_likes.post_id)                        like_count,
	IF((SELECT c.user_id FROM post_likes c WHERE user_id = ?) = ?, TRUE, FALSE) has_liked`, likes.UserId, likes.UserId).Model(&models.PostLike{}).Where("post_id = ?", likes.PostId).Error
	return postLike, err
}

func (c LikesRepository) DeleteLikes(like models.PostLike) error {
	return c.db.DB.Delete(&like).Error
}

func (c LikesRepository) GetOneLike(postId models.PostLike) (userId models.PostLike, err error) {
	return userId, c.db.DB.
		Model(&models.PostLike{}).
		Where("id = ?", postId).
		First(&postId).Error
}

package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"

	"gorm.io/gorm"
)

type PostsRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewPostsRepository(db infrastructure.Database, logger infrastructure.Logger) PostsRepository {
	return PostsRepository{
		db:     db,
		logger: logger,
	}
}

func (c PostsRepository) WithTrx(trxHandle *gorm.DB) PostsRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context.")
		return c
	}
	c.db.DB = trxHandle
	return c
}

// Save -> Posts
func (c PostsRepository) CreatePosts(Posts models.Posts) error {
	return c.db.DB.Create(&Posts).Error
}

// Update -> Posts
func (c PostsRepository) UpdatePosts(Posts models.Posts) error {
	return c.db.DB.Model(&models.Posts{}).
		Where("id = ?", Posts.PostId).
		Updates(map[string]interface{}{
			"post_id":  Posts.PostId,
			"title":    Posts.Title,
			"caption":  Posts.Caption,
			"user_id":  Posts.UserId,
			"likes":    Posts.Likes,
			"audience": Posts.Audience,
		}).Error
}

// Delete -> Posts
func (c PostsRepository) DeletePosts(ID int64) error {
	return c.db.DB.Where("id = ?", ID).
		Delete(&models.Posts{}).Error
}

// GetOneUser -> gets one post of postId
func (c PostsRepository) GetOnePost(postId int64) (Posts models.Posts, err error) {
	return Posts, c.db.DB.
		Model(&models.Posts{}).
		Where("id = ?", postId).
		First(&Posts).
		Error
}

// GetAllPosts -> Get All Posts
func (c PostsRepository) GetAllPosts(pagination utils.Pagination) ([]models.Posts, int64, error) {
	var Postss []models.Posts
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.Posts{})

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`Postss`.`name` LIKE ?", searchQuery))
	}

	err := queryBuilder.
		Find(&Postss).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return Postss, totalRows, err
}

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

// CreatePosts -> Post
func (c PostsRepository) CreatePosts(posts models.Post) error {
	return c.db.DB.Create(&posts).Error
}

// UpdatePosts -> Post
func (c PostsRepository) UpdatePosts(Posts models.Post) error {
	return c.db.DB.Model(&models.Post{}).
		Where("id = ?", Posts.ID).
		Updates(map[string]interface{}{
			"title":    Posts.Title,
			"caption":  Posts.Caption,
			"user_id":  Posts.UserId,
			"likes":    Posts.Likes,
			"audience": Posts.Audience,
		}).Error
}

// DeletePosts -> Post
func (c PostsRepository) DeletePosts(ID int64) error {
	return c.db.DB.Where("id = ?", ID).
		Delete(&models.Post{}).Error
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
	queryBuilder = queryBuilder.Model(&models.Post{})

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

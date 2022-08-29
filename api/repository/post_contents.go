package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"

	"gorm.io/gorm"
)

// PostContentsRepository -> database structure
type PostContentsRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewPostContentsRepository -> creates a new PostContents repository
func NewPostContentsRepository(db infrastructure.Database, logger infrastructure.Logger) PostContentsRepository {
	return PostContentsRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx enables repository with transaction
func (c PostContentsRepository) WithTrx(trxHandle *gorm.DB) PostContentsRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

// Save -> PostContents
func (c PostContentsRepository) CreatePostContents(PostContents []models.PostContents) error {
	return c.db.DB.Create(&PostContents).Error
}

// Update -> PostContents
func (c PostContentsRepository) UpdatePostContents(PostContents models.PostContents) error {
	return c.db.DB.Model(&models.PostContents{}).
		Where("id = ?", PostContents.ContentId).
		Updates(map[string]interface{}{
			"content_id":  PostContents.ContentId,
			"content_url": PostContents.ContentId,
			"post_id":     PostContents.PostId,
		}).Error
}

// Delete -> PostContents
func (c PostContentsRepository) DeletePostContents(ID int64) error {
	return c.db.DB.Where("id = ?", ID).
		Delete(&models.PostContents{}).Error
}

// GetAllPostContents -> Get All PostContentss
func (c PostContentsRepository) GetAllPostContentss(pagination utils.Pagination) ([]models.PostContents, int64, error) {
	var PostContentss []models.PostContents
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.PostContents{})

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`PostContents`.`name` LIKE ?", searchQuery))
	}

	err := queryBuilder.
		Find(&PostContentss).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return PostContentss, totalRows, err
}

package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"

	"gorm.io/gorm"
)

// UserRepository -> database structure
type UserRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewUserRepository -> creates a new User repository
func NewUserRepository(db infrastructure.Database, logger infrastructure.Logger) UserRepository {
	return UserRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx enables repository with transaction
func (c UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

// GetOneUser -> gets one user of userId
func (c UserRepository) GetOneUser(userId int64) (user models.User, err error) {
	return user, c.db.DB.
		Model(&models.User{}).
		Where("id = ?", userId).
		First(&user).
		Error
}

// CreateUser -> User
func (c UserRepository) CreateUser(user models.User) error {
	return c.db.DB.Create(&user).Error
}

// UpdateUser -> User
func (c UserRepository) UpdateUser(user models.User) error {
	return c.db.DB.Model(&models.User{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"email":       user.Email,
			"full_name":   user.FullName,
			"created_at":  user.CreatedAt,
			"updated_at":  user.UpdatedAt,
			"deleted_at":  user.DeletedAt,
			"profile_url": user.ProfileUrl,
			"bio":         user.Bio,
			"cover_url":   user.CoverUrl,
			"is_creator":  user.IsCreator,
		}).Error
}

// DeleteUser -> User
func (c UserRepository) DeleteUser(ID int64) error {
	return c.db.DB.Where("id = ?", ID).
		Delete(&models.User{}).Error
}

// GetAllUsers -> Get All users
func (c UserRepository) GetAllUsers(pagination utils.Pagination) ([]models.User, int64, error) {
	var users []models.User
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.User{})

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`users`.`name` LIKE ?", searchQuery))
	}

	err := queryBuilder.
		Find(&users).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return users, totalRows, err
}

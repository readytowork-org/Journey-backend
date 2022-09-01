package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"

	"gorm.io/gorm"
)

type FollowRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewFollowRepository(db infrastructure.Database, logger infrastructure.Logger) FollowRepository {
	return FollowRepository{
		db:     db,
		logger: logger,
	}
}

func (c FollowRepository) WithTrx(trxHandle *gorm.DB) FollowRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

func (c FollowRepository) GetFollowerCount(ID int64) (count int, err error) {
	return count, c.db.DB.Select("COUNT(user_id)").Where("follow_user_id = ?", ID).Find(&count).Error
}

func (c FollowRepository) GetFollowingCount(ID int64) (count int, err error) {
	return count, c.db.DB.Select("COUNT(follow_user_id)").Where("user_id = ?", ID).Find(&count).Error
}

func (c FollowRepository) GetFollowings(ID int64) (follower []models.User, err error) {
	return follower, c.db.DB.Model(&models.User{}).Select("users.*").Joins("left join followers on followers.follow_user_id = users.user_id ").Where("user_id = ?", ID).Group("users.id").Find(&follower).Error

}

func (c FollowRepository) GetFollowers(ID int64) (follower []models.User, err error) {
	return follower, c.db.DB.Model(&models.User{}).Select("users.*").Joins("left join followers on followers.user_id = users.user_id ").Where("follow_user_id = ?", ID).Group("users.id").Find(&follower).Error
}

func (c FollowRepository) Follow(follow models.Follower) error {
	return c.db.DB.Create(&follow).Error
}

func (c FollowRepository) UnFollow(ID int64) error {
	return c.db.DB.Where("id = ?", ID).
		Delete(&models.User{}).Error
}

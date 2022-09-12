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

func (c FollowRepository) GetFollowerCount(ID string) (count int, err error) {
	return count, c.db.DB.Model(&models.Follower{}).Select("COUNT(user_id)").Where("follow_user_id = ?", ID).Find(&count).Error
}

func (c FollowRepository) GetFollowingCount(ID string) (count int, err error) {

func (c FollowRepository) GetFollowings(ID string) (follower []models.User, err error) {
	return follower, c.db.DB.Model(&models.User{}).Select("users.*").Joins("left join followers on followers.follow_user_id = users.id ").Where("user_id = ?", ID).Group("users.id").Find(&follower).Error

}

func (c FollowRepository) GetFollowers(ID string) (follower []models.User, err error) {
	return follower, c.db.DB.Model(&models.User{}).Select("users.*").Joins("left join followers on followers.user_id = users.id ").Where("follow_user_id = ?", ID).Group("users.id").Find(&follower).Error
}

func (c FollowRepository) Follow(follow models.Follower) error {
	return c.db.DB.Create(&follow).Error
}

func (c FollowRepository) UnFollow(follow models.Follower) error {
	return c.db.DB.Where("user_id = ?", follow.UserId).Where("follow_user_id = ?", follow.FollowUserId).
		Delete(&follow).Error
}

func (c FollowRepository) Check(follow models.Follower) (isfollowing models.Follower, err error) {
	return isfollowing, c.db.DB.Where("user_id = ?", follow.UserId).Where("follow_user_id = ?", follow.FollowUserId).
		Find(&isfollowing).Error
}

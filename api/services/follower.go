package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"

	"gorm.io/gorm"
)

// FollowService -> struct
type FollowService struct {
	repository repository.FollowRepository
}

// NewFollowService -> creates a new Followservice
func NewFollowService(repository repository.FollowRepository) FollowService {
	return FollowService{
		repository: repository,
	}
}

// WithTrx -> enables repository with transaction
func (c FollowService) WithTrx(trxHandle *gorm.DB) FollowService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

// CreateFollow -> call to create the Follow
func (c FollowService) GetFollowerCount(ID string) (int, error) {

	return c.repository.GetFollowerCount(ID)
}

// GetAllFollow -> call to get all the Follow
func (c FollowService) GetFollowingCount(ID string) (int, error) {
	return c.repository.GetFollowingCount(ID)

}

func (c FollowService) GetFollowers(ID string) ([]models.User, error) {
	return c.repository.GetFollowers(ID)
}

//GetAllFollow -> call to get all the Follow
func (c FollowService) GetFollowings(ID string) ([]models.User, error) {

	return c.repository.GetFollowings(ID)
}

func (c FollowService) Follow(follow models.Follower) error {
	return c.repository.Follow(follow)
}

func (c FollowService) UnFollow(follow models.Follower) error {
	return c.repository.UnFollow(follow)
}

func (c FollowService) Check(follow models.Follower) (isFolloing models.Follower, err error) {
	return c.repository.Check(follow)
}

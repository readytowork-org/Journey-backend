package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"

	"gorm.io/gorm"
)

type LikesService struct {
	repository repository.LikesRepository
}

func NewLikesService(repository repository.LikesRepository) LikesService {
	return LikesService{
		repository: repository,
	}
}

func (c LikesService) WithTrx(trxHandle *gorm.DB) LikesService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

func (c LikesService) CreateLikes(like models.PostLike) error {
	return c.repository.CreateLikes(like)
}

func (c LikesService) DeleteLikes(like models.PostLike) error {
	return c.repository.DeleteLikes(like)
}

func (c LikesService) GetUserPostLikes(likes models.PostLike) (postLike models.UserPostLike, err error) {
	return c.repository.GetUserPostLikes(likes)
}

func (c LikesService) GetOneLike(postId models.PostLike) (userId models.PostLike, err error) {
	return c.repository.GetOneLike(postId)
}

// func (c LikesService) GetUsersOfPostLikes(ID int64) (userId models.User, err error) {
// 	return c.repository.GetUsersOfPostLikes(ID)
// }

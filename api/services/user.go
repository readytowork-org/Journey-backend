package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
	"boilerplate-api/utils"

	"gorm.io/gorm"
)

// UserService -> struct
type UserService struct {
	repository repository.UserRepository
}

// NewUserService -> creates a new Userservice
func NewUserService(repository repository.UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

// WithTrx -> enables repository with transaction
func (c UserService) WithTrx(trxHandle *gorm.DB) UserService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

// GetOneUser -> gets one user of userId
func (c UserService) GetOneUser(userId string) (user models.User, err error) {
	return c.repository.GetOneUser(userId)
}

// CreateUser -> call to create the User
func (c UserService) CreateUser(user models.User) error {
	err := c.repository.CreateUser(user)
	return err
}

// GetAllUser -> call to get all the User
func (c UserService) GetAllUsers(pagination utils.Pagination) ([]models.User, int64, error) {
	return c.repository.GetAllUsers(pagination)
}

func (c UserService) UpdateUser(user models.User) error {
	return c.repository.UpdateUser(user)
}

func (c UserService) DeleteUser(ID string) error {
	return c.repository.DeleteUser(ID)
}

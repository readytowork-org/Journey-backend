package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/constants"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserController -> struct
type UserController struct {
	logger      infrastructure.Logger
	userService services.UserService
	env         infrastructure.Env
}

// NewUserController -> constructor
func NewUserController(
	logger infrastructure.Logger,
	userService services.UserService,
	env infrastructure.Env,
) UserController {
	return UserController{
		logger:      logger,
		userService: userService,
		env:         env,
	}
}

// CreateUser -> Create User
func (cc UserController) CreateUser(c *gin.Context) {
	user := models.User{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&user); err != nil {
		cc.logger.Zap.Error("Error [CreateUser] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind user data")
		responses.HandleError(c, err)
		return
	}

	if err := cc.userService.WithTrx(trx).CreateUser(user); err != nil {
		cc.logger.Zap.Error("Error [CreateUser] [db CreateUser]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create user")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "User Created Sucessfully")
}

// UpdateUser -> Update User
func (cc UserController) UpdateUser(c *gin.Context) {
	user := models.User{}
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		cc.logger.Zap.Error("Error [UpdateUser] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind user data")
		responses.HandleError(c, err)
		return
	}

	_, err := cc.userService.GetOneUser(id)

	if err != nil {
		cc.logger.Zap.Error("Error [DeleteUser] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "User not found !!")
		responses.HandleError(c, err)
		return
	}
	user.ID = id
	if err := cc.userService.UpdateUser(user); err != nil {
		cc.logger.Zap.Error("Error [UpdateUser] [db UpdateUser]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Update user")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "User Updated Sucessfully")
}

// DeleteUser -> Delete User
func (cc UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := cc.userService.GetOneUser(id)

	if err != nil {
		cc.logger.Zap.Error("Error [DeleteUser] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "User not found !!")
		responses.HandleError(c, err)
		return
	}

	err = cc.userService.DeleteUser(id)

	if err != nil {
		cc.logger.Zap.Error("Error [DeleteUser] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse user ID")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "User Deleted Sucessfully")

}

// GetAllUser -> Get All User
func (cc UserController) GetAllUsers(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	users, count, err := cc.userService.GetAllUsers(pagination)

	if err != nil {
		cc.logger.Zap.Error("Error finding user records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get users data")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, users, count)
}

func (cc UserController) SearchUser(c *gin.Context) {
	queryString := c.Query("queryString")

	users, err := cc.userService.SearchUser(queryString)

	if err != nil {
		cc.logger.Zap.Error("Error [DeleteUser] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse user ID")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, users)

}

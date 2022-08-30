package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/constants"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CommentController -> struct
type FollowController struct {
	logger        infrastructure.Logger
	followService services.FollowService
	env           infrastructure.Env
}

// NewFollowController -> constructor
func NewFollowController(
	logger infrastructure.Logger,
	followService services.FollowService,
	env infrastructure.Env,
) FollowController {
	return FollowController{
		logger:        logger,
		followService: followService,
		env:           env,
	}
}

func (cc FollowController) GetFollowerCount(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		cc.logger.Zap.Error("Error [GetFollowerCount] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse  ID")
		responses.HandleError(c, err)
		return
	}
	followCount, err := cc.followService.GetFollowerCount(int64(id))

	if err != nil {
		cc.logger.Zap.Error("Error [GetFollowerCount] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get follower count ID")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, followCount)

}

func (cc FollowController) GetFollowingCount(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		cc.logger.Zap.Error("Error [GetFollowingCount] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse ID")
		responses.HandleError(c, err)
		return
	}
	followingCount, err := cc.followService.GetFollowingCount(int64(id))
	responses.SuccessJSON(c, http.StatusOK, followingCount)

	if err != nil {
		cc.logger.Zap.Error("Error [DeleteUser] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse user ID")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, followingCount)

}

func (cc FollowController) GetFollowers(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		cc.logger.Zap.Error("Error [GetFollowers] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse ID")
		responses.HandleError(c, err)
		return
	}
	followers, err := cc.followService.GetFollowers(int64(id))

	if err != nil {
		cc.logger.Zap.Error("Error getting followers", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get followers data")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, followers, 5)

}
func (cc FollowController) GetFollowings(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		cc.logger.Zap.Error("Error [GetFollowings] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse ID")
		responses.HandleError(c, err)
		return
	}
	following, err := cc.followService.GetFollowings(int64(id))

	if err != nil {
		cc.logger.Zap.Error("Error getting followings", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get followings data")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, following, 5)

}

func (cc FollowController) Follow(c *gin.Context) {

	follow := models.Follower{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&follow); err != nil {
		cc.logger.Zap.Error("Error [Follow] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind user data")
		responses.HandleError(c, err)
		return
	}

	if err := cc.followService.WithTrx(trx).Follow(follow); err != nil {
		cc.logger.Zap.Error("Error [Folloe] [db Follow]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create user")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Following successfully")

}
func (cc FollowController) UnFollow(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		cc.logger.Zap.Error("Error [GetFollowings] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse ID")
		responses.HandleError(c, err)
		return
	}
	unfollowed := cc.followService.UnFollow(int64(id))
	//todo
	responses.SuccessJSON(c, http.StatusOK, unfollowed)

}

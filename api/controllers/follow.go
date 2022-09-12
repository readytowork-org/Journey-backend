package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/constants"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

	id := c.Param("id")

	followCount, err := cc.followService.GetFollowerCount(id)

	if err != nil {
		cc.logger.Zap.Error("Error [GetFollowerCount] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get follower count ID")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, followCount)

}

func (cc FollowController) GetFollowingCount(c *gin.Context) {

	id := c.Param("id")

	followingCount, err := cc.followService.GetFollowingCount(id)
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

	id := c.Param("id")

	followers, err := cc.followService.GetFollowers(id)

	if err != nil {
		cc.logger.Zap.Error("Error getting followers", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get followers data")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, followers, 5)

}
func (cc FollowController) GetFollowings(c *gin.Context) {
	id := c.Param("id")
	following, err := cc.followService.GetFollowings(id)

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

	userId := c.Query(constants.UID)
	follow.FollowUserId = c.Param("id")
	follow.UserId = userId

	if err := cc.followService.Follow(follow); err != nil {
		if err := cc.followService.UnFollow(follow); err != nil {
			cc.logger.Zap.Error("Error [Folloe] [db Follow]: ", err.Error())
			err := errors.InternalError.Wrap(err, "Failed to create user")
			responses.HandleError(c, err)
			return
		}
		responses.SuccessJSON(c, http.StatusOK, "Un successfully")
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Foll successfully")

}

func (cc FollowController) Check(c *gin.Context) {

	follow := models.Follower{}

	userId := c.Query(constants.UID)
	follow.FollowUserId = c.Param("id")
	follow.UserId = userId

	isFollowing, err := cc.followService.Check(follow)
	if err != nil {
		cc.logger.Zap.Error("Error [Folloe] [db Follow]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create user")
		responses.HandleError(c, err)
		return

	}

	responses.JSON(c, http.StatusOK, isFollowing)

}

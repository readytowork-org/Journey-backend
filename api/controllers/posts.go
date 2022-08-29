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
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostsController struct {
	logger              infrastructure.Logger
	PostsService        services.PostsService
	PostContentsService services.PostContentsService
	env                 infrastructure.Env
}

// NewPostsController -> constructor
func NewPostsController(
	logger infrastructure.Logger,
	postsService services.PostsService,
	PostContentsService services.PostContentsService,
	env infrastructure.Env,

) PostsController {
	return PostsController{
		logger:              logger,
		PostsService:        postsService,
		PostContentsService: PostContentsService,
		env:                 env,
	}
}

// CreatePosts -> Create Posts
func (cc PostsController) CreatePosts(c *gin.Context) {
	Posts := models.Posts{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&Posts); err != nil {
		cc.logger.Zap.Error("Error [CreatePosts] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind Posts data")
		responses.HandleError(c, err)
		return
	}

	if err := cc.PostsService.WithTrx(trx).CreatePosts(Posts); err != nil {
		cc.logger.Zap.Error("Error [CreatePosts] [db CreatePosts]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create Posts")
		responses.HandleError(c, err)
		return
	}
	postContents := []models.PostContents{}

	for _, postC := range Posts.PostContents {
		postC.PostId = Posts.PostId
		postContents = append(postContents, postC)
	}

	if err := cc.PostContentsService.WithTrx(trx).CreatePostContents(postContents); err != nil {
		cc.logger.Zap.Error("Error [CreatePostContents] [db CreatePostContents] : ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create Posts")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Posts Created Sucessfully")
}

// UpdatePosts -> Update Posts
func (cc PostsController) UpdatePosts(c *gin.Context) {
	Posts := models.Posts{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&Posts); err != nil {
		cc.logger.Zap.Error("Error [UpdatePosts] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind Posts data")
		responses.HandleError(c, err)
		return
	}

	if err := cc.PostsService.WithTrx(trx).UpdatePosts(Posts); err != nil {
		cc.logger.Zap.Error("Error [UpdatePosts] [db UpdatePosts]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Update Posts")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Posts Updated Sucessfully")
}

// DeletePosts -> Delete Posts
func (cc PostsController) DeletePosts(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}

	err = cc.PostsService.DeletePosts(int64(id))

	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Posts Deleted Sucessfully")

}

// GetAllPosts -> Get All Posts
func (cc PostsController) GetAllPosts(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	Posts, count, err := cc.PostsService.GetAllPosts(pagination)

	if err != nil {
		cc.logger.Zap.Error("Error finding Posts records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get Posts data")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, Posts, count)
}

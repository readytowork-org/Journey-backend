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
	LikesService        services.LikesService
	env                 infrastructure.Env
}

// NewPostsController -> constructor
func NewPostsController(
	logger infrastructure.Logger,
	postsService services.PostsService,
	PostContentsService services.PostContentsService,
	LikesService services.LikesService,
	env infrastructure.Env,

) PostsController {
	return PostsController{
		logger:              logger,
		PostsService:        postsService,
		PostContentsService: PostContentsService,
		LikesService:        LikesService,
		env:                 env,
	}
}

// CreatePosts -> Create Post
func (cc PostsController) CreatePosts(c *gin.Context) {
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	posts := models.Post{}
	if err := c.ShouldBindJSON(&posts); err != nil {
		cc.logger.Zap.Error("Error [CreatePosts] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind Post data")
		responses.HandleError(c, err)
		return
	}

	if err := cc.PostsService.WithTrx(trx).CreatePosts(posts); err != nil {
		cc.logger.Zap.Error("Error [CreatePosts] [db CreatePosts]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create Post")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Post Created Successfully")
}

// UpdatePosts -> Update Post
func (cc PostsController) UpdatePosts(c *gin.Context) {
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	posts := models.Post{}
	if err := c.ShouldBindJSON(&posts); err != nil {
		cc.logger.Zap.Error("Error [UpdatePosts] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind Post data")
		responses.HandleError(c, err)
		return
	}

	if err := cc.PostsService.WithTrx(trx).UpdatePosts(posts); err != nil {
		cc.logger.Zap.Error("Error [UpdatePosts] [db UpdatePosts]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Update Post")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Post Updated Sucessfully")
}

func (cc PostsController) PostLikes(c *gin.Context) {
	userId := c.MustGet(constants.UID).(int64)
	id, err := strconv.Atoi(c.Param("postId"))

	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}
	posts, err := cc.PostsService.GetOnePost(int64(id), userId)

	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}
	postLike := models.PostLike{PostId: posts.ID, UserId: userId}
	err = cc.LikesService.CreateLikes(postLike)
	if err != nil {
		err = cc.LikesService.DeleteLikes(postLike)
		if err != nil {
			cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
			err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
			responses.HandleError(c, err)
			return
		}
	}
	userPostLike, err := cc.LikesService.GetUserPostLikes(postLike)

	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}

	responses.JSON(c, http.StatusOK, userPostLike)

}

// DeletePosts -> Delete Posts
func (cc PostsController) DeletePosts(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Post ID")
		responses.HandleError(c, err)
		return
	}

	err = cc.PostsService.DeletePosts(int64(id))
	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Post ID")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Post Deleted Successfully")
}

// GetAllPosts -> Get All Post
func (cc PostsController) GetAllPosts(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	Posts, count, err := cc.PostsService.GetAllPosts(pagination)

	if err != nil {
		cc.logger.Zap.Error("Error finding Post records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get Post data")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, Posts, count)
}

// GetAllPosts -> Get All Post
func (cc PostsController) GetCreatorPosts(c *gin.Context) {
	userId := c.Query(constants.UID)
	cursorPagination := utils.BuildCursorPagination(c)
	posts, err := cc.PostsService.CreatorPosts(cursorPagination, userId)

	if err != nil {
		cc.logger.Zap.Error("Error finding Post records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get Post data")
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, posts)
}

func (cc PostsController) GetUserFeeds(c *gin.Context) {
	userId := c.Query(constants.UID)
	cursorPagination := utils.BuildCursorPagination(c)
	posts, err := cc.PostsService.GetUserFeeds(cursorPagination, userId)

	if err != nil {
		cc.logger.Zap.Error("Error finding Post records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get Post data")
		responses.HandleError(c, err)
		return
	}

	responses.JSON(c, http.StatusOK, posts)
}
func (cc PostsController) GetOnePost(c *gin.Context) {
	userId := c.MustGet(constants.UID).(int64)

	id, err := strconv.Atoi(c.Param("postId"))

	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}
	posts, err := cc.PostsService.GetOnePost(int64(id), userId)

	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, posts)
}

func (cc PostsController) UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")

	if err != nil {
		cc.logger.Zap.Error("Error [Upload File] [getting file error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to file")
		responses.HandleError(c, err)
		return
	}
	cc.PostsService.UploadFile(file.Filename)

	responses.JSON(c, http.StatusOK, "file uploaded sucessfully")
}

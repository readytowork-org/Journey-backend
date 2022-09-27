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

// CommentController -> struct
type CommentController struct {
	logger         infrastructure.Logger
	commentService services.CommentService
	postService    services.PostsService
	env            infrastructure.Env
}

// NewFollowController -> constructor
func NewCommentController(
	logger infrastructure.Logger,
	commentService services.CommentService,
	postService services.PostsService,
	env infrastructure.Env,
) CommentController {
	return CommentController{
		logger:         logger,
		commentService: commentService,
		postService:    postService,

		env: env,
	}
}

func (cc CommentController) GetOneUserComment(c *gin.Context) {
	userId := c.MustGet(constants.UID).(string)

	commentId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		cc.logger.Zap.Error("Error [GetOneUserComment] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
	}

	comment, err := cc.commentService.GetOneUserComment(commentId, userId)
	if err != nil {
		cc.logger.Zap.Error("Error [GetOneUserComment] [GetOneUserComment]: ", err.Error())
		err := errors.NotFound.Wrap(err, "Cannot find comment")
		responses.HandleError(c, err)
		return
	}

	responses.JSON(c, http.StatusOK, comment)
}

// CreateFollow -> Create Follow
func (cc CommentController) CreateComment(c *gin.Context) {
	comment := models.Comment{}
	userId := c.Query(constants.UID)
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	id, err := strconv.ParseInt(c.Param("post_id"), 10, 64)

	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}

	posts, err := cc.postService.GetPost(id)
	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		cc.logger.Zap.Error("Error [CreateUser] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind user data")
		responses.HandleError(c, err)
		return
	}
	comment.PostId = posts.ID
	comment.UserId = userId
	if err := cc.commentService.WithTrx(trx).CreateComment(comment); err != nil {
		cc.logger.Zap.Error("Error [CreateComment] [db CreateComment]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create Comment")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Comment Created Sucessfully")
}

// UpdateComment -> Update Comment
func (cc CommentController) UpdateComment(c *gin.Context) {
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)
	userId := c.MustGet(constants.UID).(string)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		cc.logger.Zap.Error("Error [UpdateComment] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse user ID")
		responses.HandleError(c, err)
		return
	}

	comment, err := cc.commentService.WithTrx(trx).GetOneComment(id)
	if err != nil {
		cc.logger.Zap.Error("Error [UpdateComment] [GetOneComment]: ", err.Error())
		err := errors.NotFound.Wrap(err, "Failed to get comment")
		responses.HandleError(c, err)
		return
	}

	if comment.UserId != userId {
		cc.logger.Zap.Error("Error [UpdateComment] [User check]: ", err.Error())
		err := errors.Forbidden.Wrap(err, "Cannot edit this comment!")
		responses.HandleError(c, err)
		return
	}

	updateComment := models.Comment{}
	if err := c.ShouldBindJSON(&updateComment); err != nil {
		cc.logger.Zap.Error("Error [UpdateComment] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind Comment data")
		responses.HandleError(c, err)
		return
	}

	if err := cc.commentService.WithTrx(trx).UpdateComment(comment); err != nil {
		cc.logger.Zap.Error("Error [UpdateComment] [db UpdateComment]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Update Comment")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Comment Updated !!!")
}

// DeleteComment -> Delete Comment
func (cc CommentController) DeleteComment(c *gin.Context) {
	userId := c.Query(constants.UID)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		cc.logger.Zap.Error("Error [DeleteComment] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Comment ID")
		responses.HandleError(c, err)
		return
	}

	comment, err := cc.commentService.GetOneComment(id)
	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [GetOneComment]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get comment")
		responses.HandleError(c, err)
		return
	}

	deleteComment := models.Comment{UserId: userId, Base: models.Base{
		ID: comment.ID,
	}}
	if err = cc.commentService.DeleteComment(deleteComment); err != nil {
		cc.logger.Zap.Error("Error [DeleteComment] [DeleteComment]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Delete Comment")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Comment Deleted Successfully")
}

func (cc CommentController) CreateCommentLike(c *gin.Context) {
	userId := c.MustGet(constants.UID).(string)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}

	comment, err := cc.commentService.GetOneUserComment(id, userId)
	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}

	commentLike := models.CommentLikes{UserId: userId, CommentId: comment.ID}
	err = cc.commentService.CreateCommentLike(commentLike)
	if err != nil {
		err = cc.commentService.DeleteCommentLike(commentLike)
		if err != nil {
			cc.logger.Zap.Error("Error [DeleteCommentLike] [Conversion Error]: ", err.Error())
			err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
			responses.HandleError(c, err)
			return
		}
	}

	userCommentLike, err := cc.commentService.GetUserCommentLike(commentLike)
	if err != nil {
		cc.logger.Zap.Error("Error [UserCommentLike] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, userCommentLike)
}

func (cc CommentController) GetUserPostComment(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	pagination := utils.BuildCursorPagination(c)
	if err != nil {
		cc.logger.Zap.Error("Error [UpdateComment] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse user ID")
		responses.HandleError(c, err)
		return
	}
	post, err := cc.postService.GetPost(id)
	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}

	comments, count, err := cc.commentService.GetUserPostComment(pagination, post.ID)

	if err != nil {
		cc.logger.Zap.Error("Error finding comments for this post", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get comments for post")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, comments, count)
}

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
type CommentController struct {
	logger         infrastructure.Logger
	commentService services.CommentService
	env            infrastructure.Env
}

// NewFollowController -> constructor
func NewCommentController(
	logger infrastructure.Logger,
	commentService services.CommentService,
	env infrastructure.Env,
) CommentController {
	return CommentController{
		logger:         logger,
		commentService: commentService,
		env:            env,
	}
}

// CreateFollow -> Create Follow
func (cc CommentController) CreateComment(c *gin.Context) {
	comment := models.Comment{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&comment); err != nil {
		cc.logger.Zap.Error("Error [CreateUser] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind user data")
		responses.HandleError(c, err)
		return
	}

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
	comment := models.Comment{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := c.ShouldBindJSON(&comment); err != nil {
		cc.logger.Zap.Error("Error [UpdateComment] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind Comment data")
		responses.HandleError(c, err)
		return
	}

	comment.ID = id

	if err := cc.commentService.WithTrx(trx).UpdateComment(comment); err != nil {
		cc.logger.Zap.Error("Error [UpdateComment] [db UpdateComment]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Update Comment")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Comment Updated Sucessfully")
}

// DeleteComment -> Delete Comment
func (cc CommentController) DeleteComment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		cc.logger.Zap.Error("Error [DeleteComment] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Comment ID")
		responses.HandleError(c, err)
		return
	}

	err = cc.commentService.DeleteComment(int64(id))

	if err != nil {
		cc.logger.Zap.Error("Error [DeleteComment] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Comment ID")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Comment Deleted Sucessfully")

}

func (cc CommentController) CreateCommentLike(c *gin.Context) {
	userId := c.MustGet(constants.UID).(int64)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}

	cc.commentService.GetOneComment(int64(id), userId)

	if err != nil {
		cc.logger.Zap.Error("Error [DeletePosts] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse Posts ID")
		responses.HandleError(c, err)
		return
	}
	commentLike := models.CommentLikes{UserId: userId, CommentId: int64(id)}
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

// func (cc CommentController) DeleteCommentLike(c *gin.Context) {}

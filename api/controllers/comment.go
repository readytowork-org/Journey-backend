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

// GetAllComment -> Get All Comment
func (cc CommentController) GetAllComments(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	comments, count, err := cc.commentService.GetAllComments(pagination)

	if err != nil {
		cc.logger.Zap.Error("Error finding Comment records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get Comments data")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, comments, count)
}

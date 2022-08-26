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

// PostContentsController -> struct
type PostContentsController struct {
	logger      infrastructure.Logger
	PostContentsService services.PostContentsService
	env         infrastructure.Env
}

// NewPostContentsController -> constructor
func NewPostContentsController(
	logger infrastructure.Logger,
	PostContentsService services.PostContentsService,
	env infrastructure.Env,
) PostContentsController {
	return PostContentsController{
		logger:      logger,
		PostContentsService: PostContentsService,
		env:         env,
	}
}

// CreatePostContents -> Create PostContents
func (cc PostContentsController) CreatePostContents(c *gin.Context) {
	PostContents := models.PostContents{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&PostContents); err != nil {
		cc.logger.Zap.Error("Error [CreatePostContents] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind PostContents data")
		responses.HandleError(c, err)
		return
	}

	if err := cc.PostContentsService.WithTrx(trx).CreatePostContents(PostContents); err != nil {
		cc.logger.Zap.Error("Error [CreatePostContents] [db CreatePostContents]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create PostContents")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "PostContents Created Sucessfully")
}

// UpdatePostContents -> Update PostContents
func (cc PostContentsController) UpdatePostContents(c *gin.Context) {
	PostContents := models.PostContents{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&PostContents); err != nil {
		cc.logger.Zap.Error("Error [UpdatePostContents] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind PostContents data")
		responses.HandleError(c, err)
		return
	}

	if err := cc.PostContentsService.WithTrx(trx).UpdatePostContents(PostContents); err != nil {
		cc.logger.Zap.Error("Error [UpdatePostContents] [db UpdatePostContents]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Update PostContents")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "PostContents Updated Sucessfully")
}

// DeletePostContents -> Delete PostContents
func (cc PostContentsController) DeletePostContents(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		cc.logger.Zap.Error("Error [DeletePostContents] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse PostContents ID")
		responses.HandleError(c, err)
		return
	}

	err = cc.PostContentsService.DeletePostContents(int64(id))

	if err != nil {
		cc.logger.Zap.Error("Error [DeletePostContents] [Conversion Error]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Parse PostContents ID")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "PostContents Deleted Sucessfully")

}

// GetAllPostContents -> Get All PostContents
func (cc PostContentsController) GetAllPostContentss(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	PostContentss, count, err := cc.PostContentsService.GetAllPostContentss(pagination)

	if err != nil {
		cc.logger.Zap.Error("Error finding PostContents records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get PostContentss data")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, PostContentss, count)
}

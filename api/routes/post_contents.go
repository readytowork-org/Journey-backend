package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

// PostContentsRoutes -> struct
type PostContentsRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	PostContentsController controllers.PostContentsController
	middleware     middlewares.FirebaseAuthMiddleware
	trxMiddleware  middlewares.DBTransactionMiddleware
}

// Setup PostContents routes
func (i PostContentsRoutes) Setup() {
	i.logger.Zap.Info(" Setting up PostContents routes")
	PostContentss := i.router.Gin.Group("/PostContentss")
	{
		PostContentss.GET("", i.PostContentsController.GetAllPostContentss)
		PostContentss.POST("", i.trxMiddleware.DBTransactionHandle(), i.PostContentsController.CreatePostContents)
		PostContentss.DELETE("/:id", i.PostContentsController.DeletePostContents)
	}
}

// NewPostContentsRoutes -> creates new PostContents controller
func NewPostContentsRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	PostContentsController controllers.PostContentsController,
	middleware middlewares.FirebaseAuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
) PostContentsRoutes {
	return PostContentsRoutes{
		router:         router,
		logger:         logger,
		PostContentsController: PostContentsController,
		middleware:     middleware,
		trxMiddleware:  trxMiddleware,
	}
}

package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

// CommentRoutes -> struct
type CommentRoutes struct {
	logger            infrastructure.Logger
	router            infrastructure.Router
	CommentController controllers.CommentController
	middleware        middlewares.FirebaseAuthMiddleware
	trxMiddleware     middlewares.DBTransactionMiddleware
}

// Setup Post routes
func (i CommentRoutes) Setup() {
	i.logger.Zap.Info(" Setting up Comment routes")
	Posts := i.router.Gin.Group("/comment")
	{
		Posts.GET("", i.CommentController.GetAllComments)
		Posts.POST("", i.trxMiddleware.DBTransactionHandle(), i.CommentController.CreateComment)
		Posts.PUT("/:id", i.trxMiddleware.DBTransactionHandle(), i.CommentController.UpdateComment)
		Posts.DELETE("/:id", i.CommentController.DeleteComment)

	}
}

// NewCommentRoutes -> creates new Post controller
func NewCommentRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	CommentController controllers.CommentController,
	middleware middlewares.FirebaseAuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
) CommentRoutes {
	return CommentRoutes{
		router:         router,
		logger:         logger,
		CommentController: CommentController,
		middleware:     middleware,
		trxMiddleware:  trxMiddleware,
	}
}

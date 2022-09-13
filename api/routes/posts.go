package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

// PostRoutes -> struct
type PostRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	PostController controllers.PostsController
	middleware     middlewares.FirebaseAuthMiddleware
	trxMiddleware  middlewares.DBTransactionMiddleware
}

// Setup Post routes
func (i PostRoutes) Setup() {
	i.logger.Zap.Info(" Setting up Post routes")
	Posts := i.router.Gin.Group("/post")
	{
		Posts.GET("", i.PostController.GetAllPosts)
		Posts.POST("", i.trxMiddleware.DBTransactionHandle(), i.PostController.CreatePosts)
		Posts.PUT("/:id", i.trxMiddleware.DBTransactionHandle(), i.PostController.UpdatePosts)
		Posts.DELETE("/:id", i.PostController.DeletePosts)
		Posts.POST("/like/:postId", i.PostController.PostLikes)
		Posts.GET("/:id", i.PostController.GetOnePost)
		Posts.GET("/creator", i.PostController.GetCreatorPosts)
		Posts.GET("/feed", i.PostController.GetUserFeeds)
		Posts.POST("/upload", i.PostController.UploadFile)
		Posts.GET("/likeusers/:id", i.PostController.GetUsersOfPostLikes)
	}
}

// NewPostRoutes -> creates new Post controller
func NewPostRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	PostController controllers.PostsController,
	middleware middlewares.FirebaseAuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
) PostRoutes {
	return PostRoutes{
		router:         router,
		logger:         logger,
		PostController: PostController,
		middleware:     middleware,
		trxMiddleware:  trxMiddleware,
	}
}

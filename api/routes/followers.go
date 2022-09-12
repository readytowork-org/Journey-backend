package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

// FollowerRoutes -> struct
type FollowerRoutes struct {
	logger           infrastructure.Logger
	router           infrastructure.Router
	FollowController controllers.FollowController
	middleware       middlewares.FirebaseAuthMiddleware
	trxMiddleware    middlewares.DBTransactionMiddleware
}

// Setup Follower routes
func (i FollowerRoutes) Setup() {
	i.logger.Zap.Info(" Setting up Follower routes")
	Posts := i.router.Gin.Group("/follow")
	{
		Posts.GET("/follower/:id", i.FollowController.GetFollowerCount)
		Posts.GET("/following/:id", i.FollowController.GetFollowingCount)
		Posts.GET("/followers/:id", i.FollowController.GetFollowers)
		Posts.GET("/followings/:id", i.FollowController.GetFollowings)
		Posts.GET("/check/:id", i.FollowController.Check)
		Posts.GET("/:id", i.FollowController.Follow)
	}
}

// NewFollowerRoutes -> creates new Post controller
func NewFollowerRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	FollowController controllers.FollowController,
	middleware middlewares.FirebaseAuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
) FollowerRoutes {
	return FollowerRoutes{
		router:           router,
		logger:           logger,
		FollowController: FollowController,
		middleware:       middleware,
		trxMiddleware:    trxMiddleware,
	}
}

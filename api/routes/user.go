package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

// UserRoutes -> struct
type UserRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	userController controllers.UserController
	middleware     middlewares.FirebaseAuthMiddleware
	trxMiddleware  middlewares.DBTransactionMiddleware
}

// Setup user routes
func (i UserRoutes) Setup() {
	i.logger.Zap.Info(" Setting up user routes")
	users := i.router.Gin.Group("/user")
	{
		users.GET("", i.userController.GetAllUsers)
		users.POST("", i.trxMiddleware.DBTransactionHandle(), i.userController.CreateUser)
		users.DELETE("/:id", i.userController.DeleteUser)
		users.PUT("/:id", i.userController.UpdateUser)
		users.GET("/search", i.userController.SearchUser)
	}
}

// NewUserRoutes -> creates new user controller
func NewUserRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	userController controllers.UserController,
	middleware middlewares.FirebaseAuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
) UserRoutes {
	return UserRoutes{
		router:         router,
		logger:         logger,
		userController: userController,
		middleware:     middleware,
		trxMiddleware:  trxMiddleware,
	}
}

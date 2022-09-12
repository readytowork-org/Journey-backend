package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewPostRoutes),
	fx.Provide(NewFollowerRoutes),
	fx.Provide(NewCommentRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	userRoutes UserRoutes,
	postRoutes PostRoutes,
	followRoutes FollowerRoutes,
) Routes {
	return Routes{
		userRoutes,
		postRoutes,
		followRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}

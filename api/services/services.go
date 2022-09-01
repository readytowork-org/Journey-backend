package services

import "go.uber.org/fx"

// Module exports services present
var Module = fx.Options(

	fx.Provide(NewUserService),
	fx.Provide(NewPostsService),
	fx.Provide(NewFirebaseService),
	fx.Provide(NewPostContentsService),
	fx.Provide(NewLikesService),
)

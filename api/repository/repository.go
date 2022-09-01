package repository

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewUserRepository),
	fx.Provide(NewLikesRepository),
  fx.Provide(NewCommentRepository),
	fx.Provide(NewFollowRepository),
  fx.Provide(NewPostsRepository),
	fx.Provide(NewPostContentsRepository),
)

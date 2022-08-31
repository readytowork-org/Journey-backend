package models

import "time"

type PostLike struct {
	PostId    int64     `json:"post_id"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (m PostLike) TableName() string {
	return "post_likes"
}

type UserPostLike struct {
	PostId    int64 `json:"post_id"`
	HasLiked  bool  `json:"has_liked"`
	LikeCount int64 `json:"like_count"`
}

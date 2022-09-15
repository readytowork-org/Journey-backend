package models

type CommentLikes struct {
	UserId    string `json:"user_id"`
	CommentId int64  `json:"comment_id"`
}

func (m CommentLikes) TableName() string {
	return "comment_likes"
}

type UserCommentLike struct {
	CommentId int64 `json:"comment_id"`
	HasLiked  bool  `json:"has_liked"`
	LikeCount int64 `json:"like_count"`
}

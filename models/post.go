package models

const (
	Public    = "public"
	Private   = "private"
	Followers = "followers"
)

type Post struct {
	Base
	Title        string         `json:"title"`
	Caption      *string        `json:"caption"`
	UserId       string         `json:"user_id"`
	Audience     string         `json:"audience"`
	PostContents []PostContents `json:"post_contents"`
	User         User           `json:"user"`
}

func (p Post) TableName() string {
	return "posts"
}

type UserPost struct {
	Post
	HasLiked  bool  `json:"has_liked"`
	LikeCount int64 `json:"like_count"`
}
type FeedPost struct {
	Post
	HasLiked         bool  `json:"has_liked"`
	LikeCount        int64 `json:"like_count"`
	CommentLikeCount int64 `json:"comment_like_count"`
}

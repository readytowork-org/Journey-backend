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
	PostContents []PostContents `json:"posts_contents"`
}

func (p Post) TableName() string {
	return "posts"
}

type UserPost struct {
	Post
	HasLiked  bool  `json:"has_liked"`
	LikeCount int64 `json:"like_count"`
}

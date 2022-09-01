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
	Likes        int            `json:"likes"`
	Audience     string         `json:"audience"`
	PostContents []PostContents `json:"posts_contents"`
}

func (p Post) TableName() string {
	return "posts"
}

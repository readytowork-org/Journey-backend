package models

type Post struct {
	PostId   int64  `json:"post_id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	UserId   int64  `json:"user_id"`
	Likes    int    `json:"likes"`
	Audience string `json:"audience"`
}

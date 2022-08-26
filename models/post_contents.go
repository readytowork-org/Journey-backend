package models

type PostContent struct {
	Base
	ConentUrl string `json:"conent_url"`
	PostId    int64  `json:"post_id"`
}

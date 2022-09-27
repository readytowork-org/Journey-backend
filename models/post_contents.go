package models

type PostContents struct {
	Id          int64  `json:"id"`
	ContentUrl  string `json:"content_url"`
	ContentType string `json:"content_type"`
	Thumbnail   string `json:"thumbnail"`
	PostId      int64  `json:"post_id"`
}

func (m PostContents) TableName() string {
	return "post_contents"
}

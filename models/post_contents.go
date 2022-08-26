package models

type PostContents struct {
	ContentId  int64  `json:"content_id"`
	ContentUrl string `json:"content_url"`
	PostId     int64  `json:"post_id"`
}

func (m PostContents) TableName() string {
	return "post_contents"
}

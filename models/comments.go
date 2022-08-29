package models

type Comment struct {
	Base
	Comment    string `json:"comment"`
	PostId     int64  `json:"post_id"`
	Likes      int    `json:"likes"`
	ParentIdFk *int64 `json:"parent_id_fk"`
	UserId     int64  `json:"user_id"`
}

func (c Comment) TableName() string {
	return "comments"
}

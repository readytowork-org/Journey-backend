package models

type Comment struct {
	Base
	Comment    string `json:"comment"`
	PostId     int64  `json:"post_id"`
	ParentIdFk *int64 `json:"parent_id_fk"`
	UserId     string `json:"user_id"`
}

func (c Comment) TableName() string {
	return "comments"
}

type UserComment struct {
	Comment
	HasLiked  bool  `json:"has_liked"`
	LikeCount int64 `json:"like_count"`
}

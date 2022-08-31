package models

type Follower struct {
	UserId       string `json:"user_id"`
	FollowUserId string `json:"follow_user_id"`
}

func (c Follower) TableName() string {
	return "followers"
}

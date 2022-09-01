package models

import "time"

type Follower struct {
	UserId       string    `json:"user_id"`
	FollowUserId string    `json:"follow_user_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func (c Follower) TableName() string {
	return "followers"
}

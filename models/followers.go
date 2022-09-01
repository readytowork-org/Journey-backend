package models

import "time"

type Follower struct {
	UserId       int64     `json:"user_id"`
	FollowUserId int64     `json:"follow_user_id"`
	CreatedAt    time.Time `json:"created_at"`
}

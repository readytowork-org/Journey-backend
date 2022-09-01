package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         string         `json:"id"`
	Email      string         `json:"email"`
	FullName   string         `json:"full_name"`
	ProfileUrl *string        `json:"profile_url"`
	Bio        *string        `json:"bio"`
	CoverUrl   *string        `json:"cover_url"`
	IsCreator  bool           `json:"is_creator"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"` //add soft delete in gorm
}

// TableName gives table name of model
func (m User) TableName() string {
	return "users"
}

// ToMap convert User to map
func (m User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"email":       m.Email,
		"full_name":   m.FullName,
		"created_at":  m.CreatedAt,
		"updated_at":  m.UpdatedAt,
		"deleted_at":  m.DeletedAt,
		"profile_url": m.ProfileUrl,
		"bio":         m.Bio,
		"cover_url":   m.CoverUrl,
		"is_creator":  m.IsCreator,
	}
}

package models

type User struct {
	Base
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	// CreatedAt  time.Time `json:"created_at"`
	// UpdatedAt  time.Time `json:"updated_at"`
	// DeletedAt  time.Time `json:"deleted_at"`
	ProfileUrl string `json:"profile_url"`
	Bio        string `json:"bio"`
	CoverUrl   string `json:"cover_url"`
	IsCreator  bool   `json:"is_creator"`
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
		"created_at":  m.Base.CreatedAt,
		"updated_at":  m.Base.UpdatedAt,
		"deleted_at":  m.Base.DeletedAt,
		"profile_url": m.ProfileUrl,
		"bio":         m.Bio,
		"cover_url":   m.CoverUrl,
		"is_creator":  m.IsCreator,
	}
}

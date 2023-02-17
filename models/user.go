package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email,omitempty" gorm:"not null;uniqueIndex"`
	Password string `json:"password,omitempty" gorm:"not null"`
	Phone    string `json:"phone,omitempty" gorm:"not null"`
	DOB      string `json:"dob,omitempty" gorm:"not null"`
}

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"email":      user.Email,
		"phone":      user.Phone,
		"dob":        user.DOB,
		"created_at": user.CreatedAt, // got from gorm.Model
	}

	return json.Marshal(data)
}

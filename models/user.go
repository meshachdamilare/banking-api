package model

import (
	"encoding/json"
	"time"
)

type User struct {
	Email     string    `json:"email,omitempty" gorm:"primaryKey"`
	Password  string    `json:"password,omitempty" gorm:"not null"`
	Phone     string    `json:"phone,omitempty" gorm:"not null"`
	DOB       string    `json:"dob,omitempty" gorm:"not null"`
	Timestamp time.Time `json:"created_at"`
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
		"created_at": user.Timestamp,
	}

	return json.Marshal(data)
}

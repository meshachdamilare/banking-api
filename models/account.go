package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email         string `gorm:"uniqueIndex;not null"`
	User          *User  `gorm:"foreignKey:ID"`
	Balance       uint64 `gorm:"default:0"`
	AccountNumber string
}

func (acc *Account) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"user":           acc.User,
		"balance":        acc.Balance,
		"account_number": acc.AccountNumber,
	}
	return json.Marshal(data)
}

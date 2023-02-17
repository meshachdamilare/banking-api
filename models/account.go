package model

import "encoding/json"

type Account struct {
	Email         string `gorm:"primaryKey"`
	User          *User  `gorm:"foreignKey:Email"`
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

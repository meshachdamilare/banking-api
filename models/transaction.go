package model

import (
	"encoding/json"
	"time"
)

type Txn string

const (
	Deposit  Txn = "deposit"
	Withdraw Txn = "withdraw"
)

func (t Txn) Value() string {
	return string(t)
}

type Transaction struct {
	TxnID         string
	Type          string
	Amount        uint64
	AccountNumber string
	Timestamp     time.Time
}

func (transaction *Transaction) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"transaction_id":   transaction.TxnID,
		"transaction_type": transaction.Type,
		"amount":           transaction.Amount,
		"timestamp":        transaction.Timestamp,
	}

	return json.Marshal(data)
}

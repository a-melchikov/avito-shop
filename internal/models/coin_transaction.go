package models

import "time"

type CoinTransaction struct {
	ID              int       `json:"id"`
	FromUserID      int       `json:"from_user_id,omitempty"`
	ToUserID        int       `json:"to_user_id"`
	Amount          int       `json:"amount"`
	TransactionDate time.Time `json:"transaction_date"`
}

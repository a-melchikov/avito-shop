package models

import "time"

type Purchase struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	ProductID    int       `json:"product_id"`
	Quantity     int       `json:"quantity"`
	PurchaseDate time.Time `json:"purchase_date"`
}

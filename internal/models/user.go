package models

import "time"

type User struct {
	ID               int       `json:"id"`
	Username         string    `json:"username"`
	PasswordHash     string    `json:"-"`
	RegistrationDate time.Time `json:"registration_date"`
}

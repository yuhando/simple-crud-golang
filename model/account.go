package model

import "time"

type Accounts struct {
	ID             int       `json:"id"`
	Name           string    `json:"name" valid:"required"`
	Address        string    `json:"address" valid:"required"`
	IdentityNumber string    `json:"identity_number" valid:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

package model

import "time"

type DepositHistories struct {
	IDTransaction int       `json:"id"`
	AccountID     int       `json:"account_id"`
	Name          string    `json:"name"`
	Amount        int       `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

package model

import "time"

type Form struct {
	ID        int64     `json:"id"`
	FacaltyID int64     `json:"facalty_id"`
	Type      string    `json:"type"`
	Price     float64   `json:"price"`
	IsSpicy   bool      `json:"is_spicy"`
	CreatedAt time.Time `json:"created_at"`
}

type SubmitFormPrice struct {
	Price float64 `db:"price"`
}

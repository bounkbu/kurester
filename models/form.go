package model

import "time"

type Form struct {
	ID        int64     `json:"id"`
	FacultyID int64     `json:"faculty_id"`
	Type      string    `json:"type"`
	Price     float64   `json:"price"`
	IsSpicy   bool      `json:"is_spicy"`
	CreatedAt time.Time `json:"created_at"`
}

type SubmitFormPrice struct {
	Price float64 `db:"price"`
}

type SubmitFormResponse struct {
	RecommendedMenu   RecommendedMenu     `json:"recommended_menu"`
	NearestRestaurant []NearestRestaurant `json:"nearest_restaurant"`
}

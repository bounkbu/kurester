package model

import "time"

type Menu struct {
	RestaurantId int64     `json:"restaurant_id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Price        float64   `json:"price"`
	IsSpicy      bool      `json:"is_spicy"`
	CreatedAt    time.Time `json:"created_at"`
}

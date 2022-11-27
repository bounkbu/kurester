package model

import "time"

type Restaurant struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Latitude  float64   `json:"latitude" db:"latitude"`
	Longitude float64   `json:"longitude" db:"longitude"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type RestaurantPopularity struct {
	ID           int64     `json:"id"`
	RestaurantID int64     `json:"restaurant_id" db:"restaurant_id"`
	Popularity   int64     `json:"popularity" db:"popularity"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateOrUpdateRestaurantPopularityRequest struct {
	RestaurantID int64 `json:"restaurant_id"`
}

type NearestRestaurant struct {
	Name     string  `json:"name" db:"name"`
	Distance float64 `json:"distance" db:"distance"`
}

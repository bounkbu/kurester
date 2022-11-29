package model

import "time"

type Menu struct {
	Id           int64     `json:"id" db:"id"`
	RestaurantId int64     `json:"restaurant_id" db:"restaurant_id"`
	Name         string    `json:"name" db:"name"`
	PictureUrl   string    `json:"picture_url" db:"pictureUrl"`
	Type         string    `json:"type" db:"type"`
	Price        float64   `json:"price" db:"price"`
	IsSpicy      bool      `json:"is_spicy" db:"is_spicy"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type RecommendedMenu struct {
	Id         int64      `json:"id" db:"id"`
	Restaurant Restaurant `json:"restaurant"`
	Name       string     `json:"name" db:"name"`
	PictureUrl string     `json:"picture_url" db:"pictureUrl"`
	Type       string     `json:"type" db:"type"`
	Price      float64    `json:"price" db:"price"`
	IsSpicy    bool       `json:"is_spicy" db:"is_spicy"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
}

type MenuMinPrice struct {
	Type  string  `json:"type" db:"type"`
	Price float64 `json:"price" db:"price"`
}

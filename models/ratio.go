package model

type SpicynessRatio struct {
	Name    string  `json:"name" db:"Name"`
	Percent float64 `json:"percent" db:"Percent"`
}

type PriceRatio struct {
	Results map[string]int `json:"results"`
}

type FoodTypeRatio struct {
	Type    string  `json:"type" db:"Type"`
	Percent float64 `json:"percent" db:"Percent"`
}

type PopularityFromAverageMenuPrice struct {
	RestaurantName string  `json:"restaurant_name" db:"restaurant_name"`
	AveragePrice   float64 `json:"average_price" db:"average_price"`
	Popularity     int64   `json:"popularity" db:"popularity"`
}

type AveragePopularityFromPrice struct {
	Price             float64 `json:"price" db:"price"`
	AveragePopularity float64 `json:"average_popularity" db:"average_popularity"`
}

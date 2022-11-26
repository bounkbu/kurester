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

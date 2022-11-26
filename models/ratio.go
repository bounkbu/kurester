package model

type Ratio struct {
	Name    string  `json:"name" db:"Name"`
	Percent float64 `json:"percent" db:"Percent"`
}

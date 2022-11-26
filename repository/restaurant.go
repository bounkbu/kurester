package repository

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/jmoiron/sqlx"
)

type restaurantRepository struct {
	db *sqlx.DB
}

type RestaurantRepository interface {
	InsertRestarant(model.Restaurant) error
}

func NewRestaurantRepository(db *sqlx.DB) *restaurantRepository {
	return &restaurantRepository{
		db: db,
	}
}

func (r *restaurantRepository) InsertRestarant(restaurant model.Restaurant) error {
	_, err := r.db.Query(`
		INSERT INTO restaurant (name, latitude, longitude)
		VALUES (?, ?, ?)
	`,
		restaurant.Name,
		restaurant.Latitude,
		restaurant.Longitude,
	)
	return err
}

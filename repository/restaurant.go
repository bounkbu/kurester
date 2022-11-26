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
	SelectNearbyRestaurant(model.User) ([]model.Restaurant, error)
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

func (s *restaurantRepository) SelectNearbyRestaurant(user model.User) ([]model.Restaurant, error) {
	var restaurants []model.Restaurant
	err := s.db.Select(&restaurants, `
		SELECT name, latitude, longitude, SQRT(
			POW(69.1 * (latitude - %f), 2) +
			POW(69.1 * (%f - longitude) * COS(latitude / 57.3), 2)) AS distance
		FROM restaurant HAVING distance < 5 ORDER BY distance LIMIT 10;
	`, user.Latitude, user.Longitude)
	if err != nil {
		return restaurants, err
	}

	return restaurants, nil
}

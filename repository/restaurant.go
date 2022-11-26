package repository

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type restaurantRepository struct {
	db *sqlx.DB
}

type RestaurantRepository interface {
	InsertRestarant(model.Restaurant) error
	QueryPopularRestaurant() ([]model.Restaurant, error)
}

func NewRestaurantRepository(db *sqlx.DB) *restaurantRepository {
	return &restaurantRepository{
		db: db,
	}
}

func (r *restaurantRepository) generateLogger(functionName string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"Module":  "Repository",
		"Funtion": functionName,
	})
}

func (r *restaurantRepository) InsertRestarant(restaurant model.Restaurant) error {
	logger := r.generateLogger("InsertRestarant")
	_, err := r.db.Query(`
		INSERT INTO restaurant (name, latitude, longitude)
		VALUES (?, ?, ?)
	`,
		restaurant.Name,
		restaurant.Latitude,
		restaurant.Longitude,
	)
	if err != nil {
		logger.Error(err)
	}
	logger.Info("Inserted restaurant")
	return err
}

func (r *restaurantRepository) QueryPopularRestaurant() ([]model.Restaurant, error) {
	logger := r.generateLogger("GetPopularRestaurant")
	res := []model.Restaurant{}
	q := `
		SELECT *
		FROM restaurant
		JOIN restaurant_popularity ON
		restaurant.id = restaurant_popularity.restaurant_id
		ORDER BY restaurant_popularity.popularity DESC
		LIMIT 5
	`
	err := r.db.Select(&res, q)
	if err != nil {
		logger.Error(err)
		return res, err
	}
	logger.Info("Get popular restaurant")
	return res, nil
}

package repository

import (
	"fmt"

	model "github.com/BounkBU/kurester/models"
	"github.com/jmoiron/sqlx"
)

type restaurantRepository struct {
	db *sqlx.DB
}

type RestaurantRepository interface {
	InsertRestarant(model.Restaurant) error
	QueryPopularRestaurant() ([]model.Restaurant, error)
	InsertRestaurantPopularity(restaurantId int64) error
	UpdateRestaurantPopularity(restaurantId int64) error
	QueryRestaurantPopularity(restaurantId int64) (model.RestaurantPopularity, error)
	QueryNearestRestaurants(restaurantId int64) ([]model.NearestRestaurant, error)
	QueryRestaurantById(id int64) (model.Restaurant, error)
}

func NewRestaurantRepository(db *sqlx.DB) *restaurantRepository {
	return &restaurantRepository{
		db: db,
	}
}

func (r *restaurantRepository) InsertRestarant(restaurant model.Restaurant) error {
	logger := generateLogger("InsertRestarant")
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
		return err
	}

	logger.Info("Inserted restaurant")
	return err
}

func (r *restaurantRepository) QueryPopularRestaurant() ([]model.Restaurant, error) {
	logger := generateLogger("GetPopularRestaurant")
	res := []model.Restaurant{}
	q := `
		SELECT restaurant.id, restaurant.name, restaurant.latitude, restaurant.longitude
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

func (r *restaurantRepository) InsertRestaurantPopularity(restaurantId int64) error {
	logger := r.generateLogger("InsertRestaurantPopularity")

	_, err := r.db.Query(`
		INSERT INTO restaurant_popularity (restaurant_id)
		VALUES (?, ?, ?)
	`,
		restaurantId,
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Inserted restaurant popularity")
	return nil
}

func (r *restaurantRepository) UpdateRestaurantPopularity(restaurantId int64) error {
	logger := r.generateLogger("UpdateRestaurantPopularity")

	_, err := r.db.Query("UPDATE restaurant_popularity SET popularity = popularity + 1 WHERE restaurant_id=?", restaurantId)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Updated restaurant popularity")
	return nil
}

func (r *restaurantRepository) QueryRestaurantPopularity(restaurantId int64) (restaurantPopularity model.RestaurantPopularity, err error) {
	logger := r.generateLogger("QueryRestaurantPopularity")

	err = r.db.Get(&restaurantPopularity, "SELECT * from restaurant_popularity WHERE restaurant_id=?", restaurantId)
	if err != nil {
		logger.Error(err)
		return
	}

	return restaurantPopularity, nil
}

func (r *restaurantRepository) QueryNearestRestaurants(restaurantId int64) (nearestRestaurants []model.NearestRestaurant, err error) {
	logger := r.generateLogger("QueryNearestRestaurants")

	q := fmt.Sprintf(`
		SELECT id, name, SQRT(
			POW(69.1 * (latitude - (SELECT latitude FROM faculty WHERE id = %d)), 2) +
			POW(69.1 * ((SELECT longitude FROM faculty WHERE id = %d) - longitude) * COS(latitude / 57.3), 2)) AS distance
		FROM restaurant
		HAVING distance < 25
		ORDER BY distance
		LIMIT 5;
	`, restaurantId, restaurantId)

	err = r.db.Select(&nearestRestaurants, q)
	if err != nil {
		logger.Error(err)
		return
	}

	logger.Info("Get nearest restaurants")
	return nearestRestaurants, nil
}

func (r *restaurantRepository) QueryRestaurantById(id int64) (restaurant model.Restaurant, err error) {
	logger := r.generateLogger("QueryRestaurantById")

	err = r.db.Get(&restaurant, "SELECT * from restaurant WHERE id=?", id)
	if err != nil {
		logger.Error(err)
		return
	}

	return restaurant, nil
}

package service

import (
	"database/sql"
	"errors"

	model "github.com/BounkBU/kurester/models"
	"github.com/BounkBU/kurester/repository"
	log "github.com/sirupsen/logrus"
)

type restaurantService struct {
	restaurantRepository repository.RestaurantRepository
}

type RestaurantService interface {
	CreateNewRestaurant(model.Restaurant) error
	GetPopularRestaurant() ([]model.Restaurant, error)
	CreateOrUpdateRestaurantPopularity(restaurantId int64) error
	GetNearestRestaurants(facultyId int64) ([]model.NearestRestaurant, error)
}

func NewRestaurantService(restaurantRepository repository.RestaurantRepository) *restaurantService {
	return &restaurantService{
		restaurantRepository: restaurantRepository,
	}
}

func (s *restaurantService) CreateNewRestaurant(newRestaurant model.Restaurant) error {
	log.Info("Start creating new restaurant")
	defer log.Info("End creating new restaurant")

	err := s.restaurantRepository.InsertRestarant(newRestaurant)
	return err
}

func (s *restaurantService) GetPopularRestaurant() (restaurants []model.Restaurant, err error) {
	log.Info("Start getting popular restaurant")
	defer log.Info("End getting popular restaurant")

	restaurants, err = s.restaurantRepository.QueryPopularRestaurant()
	if err != nil {
		log.Error(err)
		return
	}

	return restaurants, nil
}

func (s *restaurantService) CreateOrUpdateRestaurantPopularity(restaurantId int64) error {
	log.Info("Start creating or updating restaurant popularity")
	defer log.Info("End creating or updating  restaurant popularity")

	restaurantPopularity, err := s.restaurantRepository.QueryRestaurantPopularity(restaurantId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("Error select search movie by tmdb_movie_id: %s", err.Error())
		return err
	}

	if errors.Is(err, sql.ErrNoRows) {
		err = s.restaurantRepository.InsertRestaurantPopularity(restaurantPopularity.RestaurantID)
		log.Errorf("Error insert search: %s", err.Error())
		return err
	}

	err = s.restaurantRepository.UpdateRestaurantPopularity(restaurantPopularity.RestaurantID)
	return err
}

func (s *restaurantService) GetNearestRestaurants(facultyId int64) (restaurants []model.NearestRestaurant, err error) {
	log.Info("Start getting nearest restaurants")
	defer log.Info("End getting nearest restaurants")

	restaurants, err = s.restaurantRepository.QueryNearestRestaurants(facultyId)
	if err != nil {
		log.Error(err)
		return restaurants, err
	}

	return restaurants, nil
}

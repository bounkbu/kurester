package service

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/BounkBU/kurester/repository"
	log "github.com/sirupsen/logrus"
)

type restaurantService struct {
	restaurantRepository repository.RestaurantRepository
}

type RestaurantService interface {
	CreateNewRestaurant(newRestaurant model.Restaurant) error
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

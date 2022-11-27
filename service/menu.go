package service

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/BounkBU/kurester/repository"
	log "github.com/sirupsen/logrus"
)

type menuService struct {
	menuRepository       repository.MenuRepository
	restaurantRepository repository.RestaurantRepository
}

type MenuService interface {
	CreateNewMenu(model.Menu) error
	GetRecommendedMenu(foodType string, spicyNess bool, price float64) (model.RecommendedMenu, error)
}

func NewMenuService(menuRepository repository.MenuRepository, restaurantRepository repository.RestaurantRepository) *menuService {
	return &menuService{
		menuRepository:       menuRepository,
		restaurantRepository: restaurantRepository,
	}
}

func (s *menuService) CreateNewMenu(newMenu model.Menu) error {
	log.Info("Start creating new menu")
	defer log.Info("End creating new menu")

	err := s.menuRepository.InsertMenu(newMenu)
	return err
}

func (s *menuService) GetRecommendedMenu(foodType string, spicyNess bool, price float64) (recommendedMenu model.RecommendedMenu, err error) {
	log.Info("Start getting recommended menu")
	defer log.Info("End getting recommended menu")
	menu, err := s.menuRepository.QueryRecommendedMenu(foodType, spicyNess, price)
	if err != nil {
		return
	}

	restaurant, err := s.restaurantRepository.QueryRestaurantById(menu.RestaurantId)
	if err != nil {
		return
	}

	recommendedMenu = model.RecommendedMenu{
		Id:         menu.Id,
		Restaurant: restaurant,
		Name:       menu.Name,
		PictureUrl: menu.PictureUrl,
		Type:       menu.Type,
		Price:      menu.Price,
		IsSpicy:    menu.IsSpicy,
		CreatedAt:  menu.CreatedAt,
	}

	return recommendedMenu, nil
}

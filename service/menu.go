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
	GetAllFoodType() ([]model.Menu, error)
	GetMenuMinPrice() (map[string]float64, error)
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
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Create new menu successfully")
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
		log.Error(err)
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

	log.Info("Get restaurant by id successfully")
	return recommendedMenu, nil
}

func (s *menuService) GetAllFoodType() (foodTypes []model.Menu, err error) {
	log.Info("Get all food type")
	defer log.Info("Return all food type")

	foodTypes, err = s.menuRepository.QueryAllFoodType()
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("Get all food type successfully")
	return foodTypes, nil
}

func (s *menuService) GetMenuMinPrice() (map[string]float64, error) {
	log.Info("Get getting min price of menu")
	defer log.Info("End getting min price of menu")
	results := make(map[string]float64)

	menuMinPrice, err := s.menuRepository.QueryMenuMinPrice()
	if err != nil {
		log.Error(err)
		return results, err
	}

	for _, v := range menuMinPrice {
		results[v.Type] = v.Price
	}

	log.Info("Get all min price of menu type successfully")
	return results, nil
}

package service

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/BounkBU/kurester/repository"
	log "github.com/sirupsen/logrus"
)

type menuService struct {
	menuRepository repository.MenuRepository
}

type MenuService interface {
	CreateNewMenu(model.Menu) error
	GetRecommendedMenu(foodType string, spicyNess bool, price float64) (model.Menu, error)
}

func NewMenuService(menuRepository repository.MenuRepository) *menuService {
	return &menuService{
		menuRepository: menuRepository,
	}
}

func (s *menuService) CreateNewMenu(newMenu model.Menu) error {
	log.Info("Start creating new menu")
	defer log.Info("End creating new menu")

	err := s.menuRepository.InsertMenu(newMenu)
	return err
}

func (s *menuService) GetRecommendedMenu(foodType string, spicyNess bool, price float64) (recommendedMenu model.Menu, err error) {
	log.Info("Start getting recommended menu")
	defer log.Info("End getting recommended menu")
	menus, err := s.menuRepository.QueryRecommendedMenu(foodType, spicyNess, price)
	if err != nil {
		return
	}

	return menus, nil
}

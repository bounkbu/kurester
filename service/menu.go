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

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
	GetSpicynessRatio() ([]model.Ratio, error)
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

func (s *menuService) GetSpicynessRatio() ([]model.Ratio, error) {
	log.Info("Start getting spicyness ratio")
	defer log.Info("End getting spicyness ratio")
	var spicynessRatios []model.Ratio

	ratios, err := s.menuRepository.SelectIsSpicyRatio()
	if err != nil {
		return spicynessRatios, err
	}

	for _, ratio := range ratios {
		var ratioName string
		if ratio.Name == "0" {
			ratioName = "Not Spicy"
		} else {
			ratioName = "Spicy"
		}
		spicynessRatio := model.Ratio{
			Name:    ratioName,
			Percent: ratio.Percent,
		}
		spicynessRatios = append(spicynessRatios, spicynessRatio)
	}

	return spicynessRatios, nil
}

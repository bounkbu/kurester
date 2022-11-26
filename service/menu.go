package service

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/BounkBU/kurester/pkg/util"
	"github.com/BounkBU/kurester/repository"
	log "github.com/sirupsen/logrus"
)

type menuService struct {
	menuRepository repository.MenuRepository
}

type MenuService interface {
	CreateNewMenu(model.Menu) error
	GetSpicynessRatio() ([]model.SpicynessRatio, error)
	GetPriceRatio() (model.PriceRatio, error)
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

func (s *menuService) GetSpicynessRatio() ([]model.SpicynessRatio, error) {
	log.Info("Start getting spicyness ratio")
	defer log.Info("End getting spicyness ratio")
	var spicynessRatios []model.SpicynessRatio

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
		spicynessRatio := model.SpicynessRatio{
			Name:    ratioName,
			Percent: ratio.Percent,
		}
		spicynessRatios = append(spicynessRatios, spicynessRatio)
	}

	return spicynessRatios, nil
}

func (s *menuService) GetPriceRatio() (model.PriceRatio, error) {
	log.Info("Start getting price ratio")
	defer log.Info("End getting price ratio")

	results := make(map[string]int)
	formPrice, err := s.menuRepository.SelectSubmitFormPrice()
	if err != nil {
		return model.PriceRatio{}, err
	}

	log.Info(formPrice)

	for _, form := range formPrice {
		priceRange := util.PriceCountingHelper(form.Price)
		results[priceRange] += 1
	}

	priceRatio := model.PriceRatio{
		Results: results,
	}

	return priceRatio, nil
}

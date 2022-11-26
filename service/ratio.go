package service

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/BounkBU/kurester/pkg/util"
	"github.com/BounkBU/kurester/repository"
	log "github.com/sirupsen/logrus"
)

type ratioService struct {
	ratioRepository repository.RatioRepository
}

type RatioService interface {
	GetSpicynessRatio() ([]model.SpicynessRatio, error)
	GetPriceRatio() (model.PriceRatio, error)
	GetFoodTypeRatio() ([]model.FoodTypeRatio, error)
}

func NewRatioService(ratioRepository repository.RatioRepository) *ratioService {
	return &ratioService{
		ratioRepository: ratioRepository,
	}
}

func (s *ratioService) GetSpicynessRatio() ([]model.SpicynessRatio, error) {
	log.Info("Start getting spicyness ratio")
	defer log.Info("End getting spicyness ratio")
	var spicynessRatios []model.SpicynessRatio

	ratios, err := s.ratioRepository.SelectIsSpicyRatio()
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

func (s *ratioService) GetPriceRatio() (model.PriceRatio, error) {
	log.Info("Start getting price ratio")
	defer log.Info("End getting price ratio")

	results := make(map[string]int)
	formPrice, err := s.ratioRepository.SelectSubmitFormPrice()
	if err != nil {
		return model.PriceRatio{}, err
	}

	for _, form := range formPrice {
		priceRange := util.PriceCountingHelper(form.Price)
		results[priceRange] += 1
	}

	priceRatio := model.PriceRatio{
		Results: results,
	}

	return priceRatio, nil
}

func (s *ratioService) GetFoodTypeRatio() ([]model.FoodTypeRatio, error) {
	log.Info("Start getting spicyness ratio")
	defer log.Info("End getting spicyness ratio")

	foodTypeRatio, err := s.ratioRepository.SelectFoodTypeRatio()
	if err != nil {
		return foodTypeRatio, err
	}

	return foodTypeRatio, nil
}

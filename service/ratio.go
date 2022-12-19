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
	GetPopularityFromAverageMenuPrice() ([]model.PopularityFromAverageMenuPrice, error)
	GetAveragePopularityFromPriceRange() (model.PopularityAndPriceRatio, error)
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

	ratios, err := s.ratioRepository.QueryIsSpicyRatio()
	if err != nil {
		log.Error(err)
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

	log.Info("Get spicyness ratio successfully")
	return spicynessRatios, nil
}

func (s *ratioService) GetPriceRatio() (model.PriceRatio, error) {
	log.Info("Start getting price ratio")
	defer log.Info("End getting price ratio")

	results := make(map[string]int)
	formPrice, err := s.ratioRepository.QuerySubmitFormPrice()
	if err != nil {
		log.Error(err)
		return model.PriceRatio{}, err
	}

	for _, form := range formPrice {
		priceRange := util.PriceCountingHelper(form.Price)
		results[priceRange] += 1
	}

	priceRatio := model.PriceRatio{
		Results: results,
	}

	log.Info("Get price ratio successfully")
	return priceRatio, nil
}

func (s *ratioService) GetFoodTypeRatio() ([]model.FoodTypeRatio, error) {
	log.Info("Start getting spicyness ratio")
	defer log.Info("End getting spicyness ratio")

	foodTypeRatio, err := s.ratioRepository.QueryFoodTypeRatio()
	if err != nil {
		log.Error(err)
		return foodTypeRatio, err
	}

	log.Info("Get food type ratio successfully")
	return foodTypeRatio, nil
}

func (s *ratioService) GetPopularityFromAverageMenuPrice() ([]model.PopularityFromAverageMenuPrice, error) {
	log.Info("Start getting popularity from average menu price ratio")
	defer log.Info("End getting popularity from average menu price ratio")

	popularity, err := s.ratioRepository.QueryPopularityFromAverageMenuPrice()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info("Get popularity from average menu price ratio successfully")
	return popularity, nil
}

func (s *ratioService) GetAveragePopularityFromPriceRange() (model.PopularityAndPriceRatio, error) {
	log.Info("Start getting average popularity from price range ratio")
	defer log.Info("End getting average popularity from price range ratio")

	results := make(map[string]model.ChartRatio)

	averagePopularity, err := s.ratioRepository.QueryAveragePopularityFromPrice()
	if err != nil {
		log.Error(err)
		return model.PopularityAndPriceRatio{}, err
	}

	for _, v := range averagePopularity {
		val, ok := results[v.Type]
		if !ok {
			results[v.Type] = model.ChartRatio{
				XAxis: []float64{v.Price},
				YAxis: []int{v.Popularity},
			}
		} else {
			prevXAxis := val.XAxis
			prevYAxis := val.YAxis
			results[v.Type] = model.ChartRatio{
				XAxis: append(prevXAxis, v.Price),
				YAxis: append(prevYAxis, v.Popularity),
			}
		}
	}

	priceRatio := model.PopularityAndPriceRatio{
		Results: results,
	}

	log.Info("Get popularity from average menu price ratio successfully")
	return priceRatio, nil
}

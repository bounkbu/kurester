package repository

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/jmoiron/sqlx"
)

type ratioRepository struct {
	db *sqlx.DB
}

type RatioRepository interface {
	QueryIsSpicyRatio() ([]model.SpicynessRatio, error)
	QuerySubmitFormPrice() ([]model.SubmitFormPrice, error)
	QueryFoodTypeRatio() ([]model.FoodTypeRatio, error)
}

func NewRatioRepository(db *sqlx.DB) *ratioRepository {
	return &ratioRepository{
		db: db,
	}
}

func (r *ratioRepository) QueryIsSpicyRatio() ([]model.SpicynessRatio, error) {
	logger := generateLogger("QueryIsSpicyRatio")

	var isSpicyRatio []model.SpicynessRatio
	err := r.db.Select(&isSpicyRatio, `
		SELECT is_spicy as 'Name', COUNT(is_spicy) / (SELECT COUNT(is_spicy) FROM form) * 100 as 'Percent'
		FROM form
		GROUP BY(is_spicy);
	`)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Info("Query is_spicy ratio")
	return isSpicyRatio, nil
}

func (r *ratioRepository) QuerySubmitFormPrice() ([]model.SubmitFormPrice, error) {
	logger := generateLogger("QuerySubmitFormPrice")

	var submitFormPrice []model.SubmitFormPrice
	err := r.db.Select(&submitFormPrice, `SELECT price FROM form ORDER BY price DESC;`)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Info("Query submit form price ratio")
	return submitFormPrice, nil
}

func (r *ratioRepository) QueryFoodTypeRatio() ([]model.FoodTypeRatio, error) {
	logger := generateLogger("QueryFoodTypeRatio")

	var foodTypeRatio []model.FoodTypeRatio
	err := r.db.Select(&foodTypeRatio, `
		SELECT type as 'Type', COUNT(type) as 'Percent'
		FROM form
		GROUP BY(type);
	`)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Info("Query food type ratio")
	return foodTypeRatio, nil
}

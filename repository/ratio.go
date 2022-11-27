package repository

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/jmoiron/sqlx"
)

type ratioRepository struct {
	db *sqlx.DB
}

type RatioRepository interface {
	SelectIsSpicyRatio() ([]model.SpicynessRatio, error)
	SelectSubmitFormPrice() ([]model.SubmitFormPrice, error)
	SelectFoodTypeRatio() ([]model.FoodTypeRatio, error)
}

func NewRatioRepository(db *sqlx.DB) *ratioRepository {
	return &ratioRepository{
		db: db,
	}
}

func (r *ratioRepository) SelectIsSpicyRatio() ([]model.SpicynessRatio, error) {
	var isSpicyRatio []model.SpicynessRatio
	err := r.db.Select(&isSpicyRatio, `
		SELECT is_spicy as 'Name', COUNT(is_spicy) / (SELECT COUNT(is_spicy) FROM form) * 100 as 'Percent'
		FROM form
		GROUP BY(is_spicy);
	`)
	if err != nil {
		return nil, err
	}

	return isSpicyRatio, nil
}

func (r *ratioRepository) SelectSubmitFormPrice() ([]model.SubmitFormPrice, error) {
	var submitFormPrice []model.SubmitFormPrice
	err := r.db.Select(&submitFormPrice, `SELECT price FROM form ORDER BY price DESC;`)
	if err != nil {
		return nil, err
	}

	return submitFormPrice, nil
}

func (r *ratioRepository) SelectFoodTypeRatio() ([]model.FoodTypeRatio, error) {
	var foodTypeRatio []model.FoodTypeRatio
	err := r.db.Select(&foodTypeRatio, `
		SELECT type as 'Type', COUNT(type) as 'Percent'
		FROM form
		GROUP BY(type);
	`)
	if err != nil {
		return nil, err
	}

	return foodTypeRatio, nil
}

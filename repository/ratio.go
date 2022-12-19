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
	QueryPopularityFromAverageMenuPrice() ([]model.PopularityFromAverageMenuPrice, error)
	QueryAveragePopularityFromPrice() ([]model.AveragePopularityFromPrice, error)
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
		SELECT is_spicy as 'Name', COUNT(is_spicy) / (SELECT COUNT(is_spicy) FROM`+"`kurester.form`"+`) * 100 as 'Percent'
		FROM`+"`kurester.form`"+`
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
	err := r.db.Select(&submitFormPrice, `SELECT price FROM`+"`kurester.form`"+`as form ORDER BY price DESC;`)
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
		FROM`+"`kurester.form`"+`as form
		GROUP BY(type);
	`)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Info("Query food type ratio")
	return foodTypeRatio, nil
}

func (r *ratioRepository) QueryPopularityFromAverageMenuPrice() ([]model.PopularityFromAverageMenuPrice, error) {
	logger := generateLogger("QueryPopularityFromAverageMenuPrice")

	var popularity []model.PopularityFromAverageMenuPrice
	err := r.db.Select(&popularity, `
		SELECT (SELECT name FROM`+"`kurester.restaurant`"+`WHERE id=restaurant_id) as restaurant_name, AVG(price) as average_price, MIN(popularity) popularity
		FROM (
			SELECT m.restaurant_id, price, popularity
			FROM`+"`kurester.menu`"+`m
			INNER JOIN`+"`kurester.restaurant_popularity`"+`rp
			ON m.restaurant_id = rp.restaurant_id
		) avg_price
		GROUP BY restaurant_id
		ORDER BY average_price;
	`)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Info("Query popularity from average menu price")
	return popularity, nil
}

func (r *ratioRepository) QueryAveragePopularityFromPrice() ([]model.AveragePopularityFromPrice, error) {
	logger := generateLogger("QueryPopularityFromAverageMenuPrice")

	var averagePopularity []model.AveragePopularityFromPrice
	err := r.db.Select(&averagePopularity, `
		SELECT type, price, popularity
		FROM (
			SELECT type, price, popularity
			FROM`+"`kurester.menu`"+`m
			INNER JOIN`+"`kurester.restaurant_popularity`"+`rp
			ON m.restaurant_id = rp.restaurant_id
		) avg_price
		GROUP BY price, type, popularity
		ORDER BY type;
	`)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Info("Query average popularity from price")
	return averagePopularity, nil
}

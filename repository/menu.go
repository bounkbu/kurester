package repository

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/jmoiron/sqlx"
)

type menuRepository struct {
	db *sqlx.DB
}

type MenuRepository interface {
	InsertMenu(model.Menu) error
	SelectIsSpicyRatio() ([]model.SpicynessRatio, error)
	SelectSubmitFormPrice() ([]model.SubmitFormPrice, error)
}

func NewMenuRepository(db *sqlx.DB) *menuRepository {
	return &menuRepository{
		db: db,
	}
}

func (r *menuRepository) InsertMenu(menu model.Menu) error {
	_, err := r.db.Query(`
		INSERT INTO menu (restaurant_id, name, type, price, is_spicy)
		VALUES (?, ?, ?, ?, ?)
	`,
		menu.RestaurantId,
		menu.Name,
		menu.Type,
		menu.Price,
		menu.IsSpicy,
	)
	return err
}

func (r *menuRepository) SelectIsSpicyRatio() ([]model.SpicynessRatio, error) {
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

func (r *menuRepository) SelectSubmitFormPrice() ([]model.SubmitFormPrice, error) {
	var submitFormPrice []model.SubmitFormPrice
	err := r.db.Select(&submitFormPrice, `SELECT price FROM form ORDER BY price DESC;`)
	if err != nil {
		return nil, err
	}

	return submitFormPrice, nil
}

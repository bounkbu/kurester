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
	SelectIsSpicyRatio() ([]model.Ratio, error)
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

func (r *menuRepository) SelectIsSpicyRatio() ([]model.Ratio, error) {
	var isSpicyRatio []model.Ratio
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

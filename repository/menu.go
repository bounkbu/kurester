package repository

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/jmoiron/sqlx"
)

type menuRepository struct {
	db *sqlx.DB
}

type MenuRepository interface {
	InsertMenu(restaurant model.Menu) error
}

func NewMenuRepository(db *sqlx.DB) *menuRepository {
	return &menuRepository{
		db: db,
	}
}

func (r *menuRepository) InsertMenu(restaurant model.Menu) error {
	_, err := r.db.Query(`
		INSERT INTO menu (restaurant_id, name, type, price, is_spicy)
		VALUES (?, ?, ?, ?, ?)
	`,
		restaurant.RestaurantId,
		restaurant.Name,
		restaurant.Type,
		restaurant.Price,
		restaurant.IsSpicy,
	)
	return err
}

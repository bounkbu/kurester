package repository

import (
	"errors"

	model "github.com/BounkBU/kurester/models"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type menuRepository struct {
	db *sqlx.DB
}

type MenuRepository interface {
	InsertMenu(model.Menu) error
	QueryRecommendedMenu(foodType string, spicyNess bool, price float64) (model.Menu, error)
}

func NewMenuRepository(db *sqlx.DB) *menuRepository {
	return &menuRepository{
		db: db,
	}
}

var ErrFoundMoreThanOne error = errors.New("found more than one row in db")
var ErrNotFound error = errors.New("not found in db")

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

func (r *menuRepository) QueryRecommendedMenu(foodType string, spicyNess bool, price float64) (recommendedMenu model.Menu, err error) {
	var menus []model.Menu
	err = r.db.Select(&menus, `
		SELECT *
		FROM menu
		WHERE type = ?
		AND is_spicy = ?
		HAVING price <= ?
		LIMIT 1;
	`, foodType, spicyNess, price)
	log.Info(menus)
	if err != nil {
		log.Error(err)
		return recommendedMenu, err
	}

	menuLength := len(menus)
	if menuLength == 0 {
		return recommendedMenu, ErrNotFound
	} else if menuLength > 1 {
		return recommendedMenu, ErrFoundMoreThanOne
	}

	return menus[0], nil
}

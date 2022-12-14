package repository

import (
	"errors"

	model "github.com/BounkBU/kurester/models"
	"github.com/jmoiron/sqlx"
)

type menuRepository struct {
	db *sqlx.DB
}

type MenuRepository interface {
	InsertMenu(model.Menu) error
	QueryRecommendedMenu(foodType string, spicyNess bool, price float64) (model.Menu, error)
	QueryAllFoodType() ([]model.Menu, error)
	QueryMenuMinPrice() ([]model.MenuMinPrice, error)
}

func NewMenuRepository(db *sqlx.DB) *menuRepository {
	return &menuRepository{
		db: db,
	}
}

var ErrFoundMoreThanOne error = errors.New("found more than one row in db")
var ErrNotFound error = errors.New("not found in db")

func (r *menuRepository) InsertMenu(menu model.Menu) error {
	logger := generateLogger("InsertMenu")

	_, err := r.db.Query(`
		INSERT INTO`+"`kurester.menu`"+`(restaurant_id, name, type, price, is_spicy)
		VALUES (?, ?, ?, ?, ?)
	`,
		menu.RestaurantId,
		menu.Name,
		menu.Type,
		menu.Price,
		menu.IsSpicy,
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Insert new menu")
	return nil
}

func (r *menuRepository) QueryRecommendedMenu(foodType string, spicyNess bool, price float64) (recommendedMenu model.Menu, err error) {
	logger := generateLogger("QueryRecommendedMenu")

	var menus []model.Menu
	err = r.db.Select(&menus, `
		SELECT menu.id, menu.restaurant_id, name, pictureUrl, type, price, is_spicy, created_at
		FROM`+"`kurester.menu`"+`as menu,`+"`kurester.restaurant_popularity`"+`as restaurant_popularity
		WHERE type = ?
		AND menu.restaurant_id = restaurant_popularity.restaurant_id
		AND is_spicy = ?
		HAVING price <= ?
		ORDER BY popularity DESC, price DESC;
	`, foodType, spicyNess, price)
	if err != nil {
		logger.Error(err)
		return recommendedMenu, err
	}

	menuLength := len(menus)
	if menuLength == 0 {
		logger.Error(ErrNotFound)
		return recommendedMenu, ErrNotFound
	}

	return menus[0], nil
}

func (r *menuRepository) QueryAllFoodType() (foodTypes []model.Menu, err error) {
	logger := generateLogger("QueryAllFoodType")

	q := "SELECT DISTINCT type FROM `kurester.menu`"
	err = r.db.Select(&foodTypes, q)
	if err != nil {
		logger.Error(err)
		return foodTypes, err
	}

	logger.Info("Get all food type")
	return foodTypes, nil
}

func (r *menuRepository) QueryMenuMinPrice() (menuMinPrice []model.MenuMinPrice, err error) {
	logger := generateLogger("QueryMenuMinPrice")

	q := `
		SELECT type, min(price) as price
		FROM` + "`kurester.menu`" + `
		GROUP BY type
	`
	err = r.db.Select(&menuMinPrice, q)
	if err != nil {
		logger.Error(err)
		return menuMinPrice, err
	}

	logger.Info("Get min price of menu type")
	return menuMinPrice, nil
}

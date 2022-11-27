package repository

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/jmoiron/sqlx"
)

type formRepository struct {
	db *sqlx.DB
}

type FormRepository interface {
	InsertForm(model.Form) error
}

func NewFormRepository(db *sqlx.DB) *formRepository {
	return &formRepository{
		db: db,
	}
}

func (r *formRepository) InsertForm(form model.Form) error {
	_, err := r.db.Query(`
		INSERT INTO form (faculty_id, type, price, is_spicy)
		VALUES (?, ?, ?, ?)
	`,
		form.FacaltyID,
		form.Type,
		form.Price,
		form.IsSpicy,
	)
	return err
}

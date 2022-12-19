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
	logger := generateLogger("InsertForm")

	_, err := r.db.Query(`
		INSERT INTO`+"`kurester.form`"+`(faculty_id, type, price, is_spicy)
		VALUES (?, ?, ?, ?)
	`,
		form.FacultyID,
		form.Type,
		form.Price,
		form.IsSpicy,
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Insert new form")
	return err
}

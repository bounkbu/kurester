package repository

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/jmoiron/sqlx"
)

type facultyRepository struct {
	db *sqlx.DB
}

type FacultyRepository interface {
	QueryAllFaculty() ([]model.Faculty, error)
}

func NewFacultyRepository(db *sqlx.DB) *facultyRepository {
	return &facultyRepository{
		db: db,
	}
}

func (f *facultyRepository) QueryAllFaculty() (res []model.Faculty, err error) {
	logger := generateLogger("QueryAllFaculty")
	q := `
		SELECT id, name
		FROM faculty
	`
	err = f.db.Select(&res, q)
	if err != nil {
		logger.Error(err)
		return res, err
	}
	logger.Info("Get popular restaurant")
	return res, nil
}

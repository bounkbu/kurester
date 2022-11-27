package service

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/BounkBU/kurester/repository"
	log "github.com/sirupsen/logrus"
)

type facultyService struct {
	facultyRepository repository.FacultyRepository
}

type FacultyService interface {
	GetAllFaculty() ([]model.Faculty, error)
}

func NewFacultyService(facultyRepository repository.FacultyRepository) *facultyService {
	return &facultyService{
		facultyRepository: facultyRepository,
	}
}

func (f *facultyService) GetAllFaculty() (faculties []model.Faculty, err error) {
	log.Info("Get all faculty name, id")
	defer log.Info("Return all faculty")

	faculties, err = f.facultyRepository.QueryAllFaculty()
	if err != nil {
		log.Error(err)
	}
	return
}

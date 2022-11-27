package service

import (
	model "github.com/BounkBU/kurester/models"
	"github.com/BounkBU/kurester/repository"
	log "github.com/sirupsen/logrus"
)

type formService struct {
	formRepository repository.FormRepository
}

type FormService interface {
	CreateNewForm(model.Form) error
}

func NewFormService(formRepository repository.FormRepository) *formService {
	return &formService{
		formRepository: formRepository,
	}
}

func (s *formService) CreateNewForm(newForm model.Form) error {
	log.Info("Start creating new form")
	defer log.Info("End creating new form")

	err := s.formRepository.InsertForm(newForm)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Create new form successfully")
	return err
}

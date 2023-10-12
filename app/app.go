package app

import (
	"github.com/robertobouses/http-framework_ejercicio8/entity"
	"github.com/robertobouses/http-framework_ejercicio8/repository"
)

type APP interface {
	CreateMeasurement(newMedicion entity.Measurement) error
}

type service struct {
	repo repository.REPOSITORY
}

func NewAPP(repo repository.REPOSITORY) APP {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateMeasurement(newMedicion entity.Measurement) error {
	return s.repo.InsertMeasurement(newMedicion)

}

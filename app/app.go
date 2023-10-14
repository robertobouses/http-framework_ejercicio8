package app

import (
	"log"

	"github.com/robertobouses/http-framework_ejercicio8/entity"
	"github.com/robertobouses/http-framework_ejercicio8/repository"
)

type APP interface {
	CreateMeasurement(newMedicion entity.Measurement) error
	PrintMeasurement() ([]entity.Measurement, error)
	DeleteMeasurement(id string) error
	CalcCubic(id string) (int, error)
	FindMinAndMaxCubic(cubics []int) (int, int)
	DeleteAllMeasurement() error
	DeleteEmptyMeasurement() error
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

func (s *service) PrintMeasurement() ([]entity.Measurement, error) {
	measurement, err := s.repo.ExtractMeasurement()
	if err != nil {
		log.Printf("Error al extraer mediciones", err)
		return []entity.Measurement{}, err
	}
	return measurement, nil

}

func (s *service) DeleteMeasurement(id string) error {
	err := s.repo.DeleteMeasurement(id)
	if err != nil {
		log.Printf("Error", err)
		return err
	}
	return nil
}

func (s *service) CalcCubic(id string) (int, error) {
	measurement, err := s.repo.ExtractMeasurementId(id)
	if err != nil {
		log.Printf("Error al obtener la medición por ID", err)
		return 0, err
	}
	cubic := measurement.ValorX * measurement.ValorY * measurement.ValorZ
	return cubic, nil
}

func (s *service) FindMinAndMaxCubic(cubics []int) (int, int) {
	if len(cubics) == 0 {
		return 0, 0
	}

	minCubic := cubics[0]
	maxCubic := cubics[0]

	for _, cubic := range cubics {
		if cubic < minCubic {
			minCubic = cubic
		}
		if cubic > maxCubic {
			maxCubic = cubic
		}
	}

	return minCubic, maxCubic
}

func (s *service) DeleteAllMeasurement() error {
	err := s.repo.DeleteAllMeasurement()
	if err != nil {
		log.Print("Error", err)
		return err
	}
	return nil
}

func (s *service) DeleteEmptyMeasurement() error {
	err := s.repo.DeleteAllMeasurement()
	if err != nil {
		log.Print("Error", err)
		return err
	}
	return nil
}

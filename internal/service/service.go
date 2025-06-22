package service

import (
	"echo_start/internal/model"
	"echo_start/internal/repository"
	"errors"
)

type CarService interface {
	CreateCar(car *model.Car) error
	GetCars() ([]model.Car, error)
}

type carService struct {
	repo repository.CarRepository
}

func NewServiceCar(r repository.CarRepository) CarService {
	return &carService{repo: r}
}

func (s *carService) CreateCar(car *model.Car) error {
	if car.Make == "" || car.Model == "" {
		return errors.New("Make and Model are required")
	}

	return s.repo.CreateCar(car)
}

func (s *carService) GetCars() ([]model.Car, error) {
	return s.repo.GetCars()
}

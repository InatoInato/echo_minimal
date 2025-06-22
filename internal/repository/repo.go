package repository

import (
	"echo_start/internal/config"
	"echo_start/internal/model"
)

type CarRepository interface {
	CreateCar(car *model.Car) error
	GetCars() ([]model.Car, error)
}

type carRepo struct{}

func NewCarRepository() CarRepository {
	return &carRepo{}
}

func (r *carRepo) CreateCar(car *model.Car) error {
	return config.DB.Create(car).Error
}

func (r *carRepo) GetCars() ([]model.Car, error) {
	var cars []model.Car
	err := config.DB.Find(&cars).Error
	return cars, err
}

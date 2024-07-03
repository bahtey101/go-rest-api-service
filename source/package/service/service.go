package service

import (
	"github.com/bahtey101/go-rest-api-service/model"
	"github.com/bahtey101/go-rest-api-service/package/repository"
)

type Service struct {
	repos repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{repos: *repos}
}

func (service *Service) Create(car model.Car) (*model.Car, error) {
	return service.repos.Create(car)
}

func (service *Service) GetAll() (*[]model.Car, error) {
	return service.repos.GetAll()
}

func (service *Service) Get(car model.Car) (*model.Car, error) {
	return service.repos.Get(car)
}

func (service *Service) Replace(car model.Car) (*model.Car, error) {
	return service.repos.Replace(car)
}

func (service *Service) Update(updated_car model.Car) error {
	car, err := service.Get(model.Car{ID: updated_car.ID})
	if err != nil {
		return err
	}

	if updated_car.Brand != "-" {
		car.Brand = updated_car.Brand
	}
	if updated_car.Model != "-" {
		car.Model = updated_car.Model
	}
	if updated_car.Mileage != -1 {
		car.Mileage = updated_car.Mileage
	}
	if updated_car.OwnersNumber != -1 {
		car.OwnersNumber = updated_car.OwnersNumber
	}

	_, err = service.repos.Replace(*car)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) Delete(car model.Car) error {
	return service.repos.Delete(car)
}

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
	// получить Id
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

func (service *Service) Update(car model.Car) error {
	return service.repos.Update(car)
}

func (service *Service) Delete(car model.Car) error {
	return service.repos.Delete(car)
}

package repository

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bahtey101/go-rest-api-service/model"
)

type Repository struct {
	path string
}

type FormatJSON struct {
	LastID int64       `json:"last_id"`
	Cars   []model.Car `json:"Cars"`
}

func NewRepository(path string) *Repository {
	return &Repository{path: path}
}

func (repos *Repository) Create(car model.Car) (*model.Car, error) {
	jsonInfo, err := repos.ReadJSON()
	if err != nil {
		return nil, err
	}

	jsonInfo.LastID += 1
	car.ID = jsonInfo.LastID
	jsonInfo.Cars = append(jsonInfo.Cars, car)

	err = repos.WriteJSON(*jsonInfo)
	if err != nil {
		return nil, err
	}

	return &car, nil
}

func (repos *Repository) GetAll() (*[]model.Car, error) {
	jsonInfo, err := repos.ReadJSON()
	if err != nil {
		return nil, err
	}

	return &jsonInfo.Cars, nil
}

func (repos *Repository) Get(car model.Car) (*model.Car, error) {
	jsonInfo, err := repos.ReadJSON()
	if err != nil {
		return nil, err
	}

	for _, _car := range jsonInfo.Cars {
		if _car.ID == car.ID {
			return &_car, nil
		}
	}

	return nil, errors.New("car not found")
}

func (repos *Repository) Replace(car model.Car) (*model.Car, error) {
	jsonInfo, err := repos.ReadJSON()
	if err != nil {
		return nil, err
	}

	for index, _car := range jsonInfo.Cars {
		if _car.ID == car.ID {
			jsonInfo.Cars[index] = car
		}
	}

	err = repos.WriteJSON(*jsonInfo)
	if err != nil {
		return nil, err
	}

	return &car, nil
}

func (repos *Repository) Update(car model.Car) error {
	jsonInfo, err := repos.ReadJSON()
	if err != nil {
		return err
	}

	for _, _car := range jsonInfo.Cars {
		if _car.ID == car.ID {

		}
	}

	err = repos.WriteJSON(*jsonInfo)
	if err != nil {
		return err
	}

	return nil
}

func (repos *Repository) Delete(car model.Car) error {
	jsonInfo, err := repos.ReadJSON()
	if err != nil {
		return err
	}

	for index, _car := range jsonInfo.Cars {
		if _car.ID == car.ID {
			jsonInfo.Cars = append(jsonInfo.Cars[:index], jsonInfo.Cars[index+1:]...)
		}
	}

	err = repos.WriteJSON(*jsonInfo)
	if err != nil {
		return err
	}

	return nil
}

func (repos *Repository) WriteJSON(data FormatJSON) error {
	if _, err := os.Stat(repos.path); os.IsNotExist(err) {
		data.LastID = 0
	}

	jsonInfo, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(repos.path, jsonInfo, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (repos *Repository) ReadJSON() (*FormatJSON, error) {
	data, err := os.ReadFile(repos.path)
	if err != nil {
		return nil, err
	}

	var jsonInfo FormatJSON
	err = json.Unmarshal(data, &jsonInfo)
	if err != nil {
		return nil, err
	}

	return &jsonInfo, nil
}

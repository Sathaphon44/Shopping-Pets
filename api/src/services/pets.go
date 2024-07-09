package services

import (
	"shoppets-api/domain/entities"
	"shoppets-api/domain/repository"
)

type TPetsServices struct {
	RepoPets repository.IPetsRepo
}

type IPetsServices interface {
	GetPetsAllServices() (*[]entities.Pets, error)
}

func NewPetsServiecs(repo repository.IPetsRepo) IPetsServices {
	return &TPetsServices{
		RepoPets: repo,
	}
}

func (ps TPetsServices) GetPetsAllServices() (*[]entities.Pets, error) {
	data, err := ps.RepoPets.GetPetsAll()
	if err != nil {
		return nil, err
	}
	return data, err
}

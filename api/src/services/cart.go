package services

import (
	"shoppets-api/domain/entities"
	"shoppets-api/domain/repository"
)

type TCartServices struct {
	RepoCart repository.ICartRepo
}

type ICartServices interface {
	UpdateCartServices(userId string, cart *entities.Cart) (interface{}, error)
	GetDetailCartServices(userId string) (interface{}, error)
	DeleteCartServices(userId string) (interface{}, error)
}

func NewCartServices(repo repository.ICartRepo) ICartServices {
	return &TCartServices{
		RepoCart: repo,
	}
}

func (ps TCartServices) UpdateCartServices(userId string, cart *entities.Cart) (interface{}, error) {
	var intent string
	statusUpdate, err := ps.RepoCart.UpdatePetsInCarts(userId, cart)
	if err != nil {
		return nil, err
	}
	intent = "update"
	if statusUpdate.MatchedCount == 0 {
		cart.UserId = userId
		_, err := ps.RepoCart.InsertPetsInCarts(cart)
		if err != nil {
			return nil, err
		}
		intent = "insert"
	}

	if statusUpdate.ModifiedCount != 0 || intent == "insert" {
		FindResult, err := ps.RepoCart.FindOneCart(userId)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"intent": intent,
			"result": FindResult,
		}, nil
	}
	return map[string]interface{}{
		"intent": intent,
		"result": nil,
	}, nil
}

func (ps TCartServices) GetDetailCartServices(userId string) (interface{}, error) {

	FindResult, err := ps.RepoCart.FindOneCart(userId)
	if err != nil {
		return nil, err
	}
	return FindResult, nil

}

func (ps TCartServices) DeleteCartServices(userId string) (interface{}, error) {
	result, err := ps.RepoCart.DeleteCart(userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

package repository

import (
	"context"
	. "shoppets-api/domain/datasources"
	"shoppets-api/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TPetsRepo struct {
	Context    context.Context
	Collection mongo.Collection
}

type IPetsRepo interface {
	GetPetsAll() (*[]entities.Pets, error)
}

func NewPetsRepo(db *Mongodb) *TPetsRepo {
	col := db.Client.Database("shop_pets").Collection("pets")
	return &TPetsRepo{
		Context:    db.Context,
		Collection: *col,
	}
}

func (col TPetsRepo) GetPetsAll() (*[]entities.Pets, error) {
	data, err := col.Collection.Find(col.Context, bson.M{})
	if err != nil {
		return nil, err
	}

	var pets []entities.Pets // Replace YourPetType with the actual type of your documents

	if err := data.All(col.Context, &pets); err != nil {
		return nil, err
	}

	return &pets, nil
}

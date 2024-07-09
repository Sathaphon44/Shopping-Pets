package repository

import (
	"context"
	. "shoppets-api/domain/datasources"
	"shoppets-api/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TCartRepo struct {
	Context    context.Context
	Collection mongo.Collection
}

type ICartRepo interface {
	UpdatePetsInCarts(userId string, cart *entities.Cart) (*mongo.UpdateResult, error)
	InsertPetsInCarts(cart *entities.Cart) (*mongo.InsertOneResult, error)
	FindOneCart(userId string) (interface{}, error)
	DeleteCart(userId string) (interface{}, error)
}

func NewCartRepo(db *Mongodb) *TCartRepo {
	col := db.Client.Database("shop_pets").Collection("cart")
	return &TCartRepo{
		Context:    db.Context,
		Collection: *col,
	}
}

func (col TCartRepo) UpdatePetsInCarts(userId string, cart *entities.Cart) (*mongo.UpdateResult, error) {
	filter := bson.M{"userId": userId}
	update := bson.M{
		"$set": bson.M{
			"pets": cart.Pets,
		},
	}
	updateResult, err := col.Collection.UpdateOne(
		col.Context,
		filter,
		update,
	)
	if err != nil {
		return nil, err
	}

	return updateResult, nil

}

func (col TCartRepo) InsertPetsInCarts(cart *entities.Cart) (*mongo.InsertOneResult, error) {

	insertResult, err := col.Collection.InsertOne(
		col.Context,
		cart,
	)
	if err != nil {
		return nil, err
	}

	return insertResult, nil

}

func (col TCartRepo) FindOneCart(userId string) (interface{}, error) {
	data := new(entities.DetailCart)
	filters := bson.M{"userId": bson.M{"$eq": userId}}
	if err := col.Collection.FindOne(col.Context, filters).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (col TCartRepo) DeleteCart(userId string) (interface{}, error) {
	filter := bson.M{"userId": userId}
	result, err := col.Collection.DeleteOne(col.Context, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

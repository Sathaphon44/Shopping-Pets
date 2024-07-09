package repository

import (
	"context"
	. "shoppets-api/domain/datasources"
	"shoppets-api/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	Context    context.Context
	Collection mongo.Collection
}

type IUsersRepo interface {
	FindUsers() (interface{}, error)
	CountUsers(email string) (interface{}, error)
	InsertOneUser(username, email string, password []byte) (interface{}, error)
	FindOneUser(email string) (*entities.ResponseUsers, error)
	DeleteOneUser(email string) (interface{}, error)
}

func NewUserRepo(db *Mongodb) *Users {
	col := db.Client.Database("shop_pets").Collection("users")
	return &Users{
		Context:    db.Context,
		Collection: *col,
	}
}

func (col Users) FindUsers() (interface{}, error) {
	data, err := col.Collection.Find(col.Context, bson.D{})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (col Users) FindOneUser(email string) (*entities.ResponseUsers, error) {
	var result entities.ResponseUsers
	err := col.Collection.FindOne(col.Context, bson.M{"email": email}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (col Users) InsertOneUser(username, email string, password []byte) (interface{}, error) {
	status, err := col.Collection.InsertOne(col.Context, bson.M{"username": username, "email": email, "password": password})
	if err != nil {
		return err, nil
	}
	return status, nil
}

func (col Users) CountUsers(email string) (interface{}, error) {
	filter := bson.M{"email": email}
	data, err := col.Collection.CountDocuments(col.Context, filter)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (col Users) DeleteOneUser(email string) (interface{}, error) {
	status, err := col.Collection.DeleteOne(col.Context, bson.M{"email": email})
	if err != nil {
		return nil, err
	}
	return status, nil
}

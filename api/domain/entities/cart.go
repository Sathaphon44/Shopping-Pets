package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type PetsInCart struct {
	PetId    string  `bson:"petId"`
	Name     string  `json:"name" bson:"name"`
	Age      int64   `json:"age" bson:"age"`
	Price    float64 `json:"price" bson:"price"`
	Details  string  `json:"details" bson:"details"`
	Amount   float64 `json:"amount" bson:"amount"`
	Category string  `json:"category" bson:"category"`
	Image    string  `json:"image" bson:"image"`
}

type Cart struct {
	UserId    string        `json:"userId" bson:"userId"`
	Pets      *[]PetsInCart `json:"pets" bson:"pets"`
	Timestamp string        `json:"timestamp" bson:"timestamp"`
}

type DetailCart struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    string             `json:"userId" bson:"userId"`
	Pets      *[]PetsInCart      `json:"pets" bson:"pets"`
	Timestamp string             `json:"timestamp" bson:"timestamp"`
}

type UserInCart struct {
	UserId   string `json:"userId" bson:"userId"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
}

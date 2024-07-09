package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseUsers struct {
	Uid      primitive.ObjectID `bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

type RegistrationUsers struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

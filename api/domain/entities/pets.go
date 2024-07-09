package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pets struct {
	Id      primitive.ObjectID `bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Age     int64              `json:"age" bson:"age"`
	Price   float64            `json:"price" bson:"price"`
	Details string             `json:"details" bson:"details"`
	Image   string             `json:"image" bson:"image"`
}

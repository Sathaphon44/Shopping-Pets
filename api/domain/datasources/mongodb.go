package datasources

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	Context context.Context
	Client  *mongo.Client
}

type IMongoDB interface {
}

func ConnectMongoDB() *Mongodb {
	uri := os.Getenv("MONGDB_URL")
	ctx := context.Background()
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	return &Mongodb{
		Context: ctx,
		Client:  client,
	}
}

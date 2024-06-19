package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

func InitDB(url string) (*Database, error) {

	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err != nil {
		cancel()
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		cancel()
		return nil, err
	}
	fmt.Println("Connected to MongoDB")
	return &Database{
		Client: client,
	}, nil
}

package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	UpdateOne(ctx context.Context, filer interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

type Database struct {
	Client *mongo.Collection
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
		Client: client.Database("gslab").Collection("employee"),
	}, nil
}

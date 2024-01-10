package repo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client *mongo.Client
}

func ConnectMongoDB() *MongoClient {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@db:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")

	return &MongoClient{client: client}
}

func (mongo *MongoClient) NewMongoDB(db string) *mongo.Database {
	return mongo.client.Database(db)
}

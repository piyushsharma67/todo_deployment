package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func Init(connString string, dbName string) (*Database,error) {
	
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))

	if err != nil {
		return nil,err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil,fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	fmt.Println("Connected to db!!")

	return &Database{
		Client: client,
		Database: client.Database(dbName),
	},nil

}

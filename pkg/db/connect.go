package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/PoulDev/roommates-api/config"

	"time"
	"context"
)


// Internal Variables
var (
	client *mongo.Client
	database *mongo.Database

	users *mongo.Collection
	carts *mongo.Collection
	houses *mongo.Collection
)

func ConnectDatabase() error {
	credentials := options.Credential{
		Username: config.DBUser,
		Password: config.DBPassword,
		AuthSource: config.DBName,
	}
	clientOptions := options.Client().
					SetAuth(credentials).
					SetHosts([]string{config.DBHost})


	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error;
	client, err = mongo.Connect(ctx, clientOptions)
	if (err != nil) {
		return err;
	}

	if err := CheckConnection(); err != nil {
		return err;
	}

	database = client.Database(config.DBName)
	users = database.Collection("users")
	carts = database.Collection("carts")
	houses = database.Collection("houses")

	return nil;
}

func CheckConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
    
	return client.Ping(ctx, nil);
}

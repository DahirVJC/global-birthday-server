package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func getMongoUrl() string {
	dbURL := os.Getenv("MONGODB_URI")

	docs := "www.mongodb.com/docs/drivers/go/current/"
	if dbURL == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}

	return dbURL
}

func Connect() (*mongo.Client, error) {
	dbURL := getMongoUrl()

	client, err := mongo.Connect(options.Client().
		ApplyURI(dbURL))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

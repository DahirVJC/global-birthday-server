package repositories

import (
	"global-birthday-server/mongodb"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func NewUsersController() *UserRepository {
	mongoClient, _ := mongodb.Connect()
	return &UserRepository{
		client:     mongoClient,
		collection: mongoClient.Database(os.Getenv("MONGODB_NAME")).Collection("users"),
	}
}

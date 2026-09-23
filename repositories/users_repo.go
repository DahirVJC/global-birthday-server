package repositories

import (
	"global-birthday-server/models"
	"global-birthday-server/mongodb"
	"os"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DbUserRepository struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func NewUsersController() *DbUserRepository {
	mongoClient, _ := mongodb.Connect()
	return &DbUserRepository{
		client:     mongoClient,
		collection: mongoClient.Database(os.Getenv("MONGODB_NAME")).Collection("users"),
	}
}

func (ur DbUserRepository) Get() ([]*models.User, error) {
	return nil, nil
}

func (ur DbUserRepository) GetById(id uuid.UUID) (*models.User, error) {
	return nil, nil
}

func (ur DbUserRepository) Create(user models.User) error {
	return nil
}

func (ur DbUserRepository) Update(id uuid.UUID, user models.User) error {
	return nil
}

func (ur DbUserRepository) Delete(id uuid.UUID) error {
	return nil
}

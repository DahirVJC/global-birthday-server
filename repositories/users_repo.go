package repositories

import (
	"global-birthday-server/models"
	"global-birthday-server/mongodb"
	"os"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DbUsersRepository struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func NewUsersController() *DbUsersRepository {
	mongoClient, _ := mongodb.Connect()
	return &DbUsersRepository{
		client:     mongoClient,
		collection: mongoClient.Database(os.Getenv("MONGODB_NAME")).Collection("users"),
	}
}

func (ur DbUsersRepository) Get() ([]*models.User, error) {
	return nil, nil
}

func (ur DbUsersRepository) GetById(id uuid.UUID) (*models.User, error) {
	return nil, nil
}

func (ur DbUsersRepository) Create(user models.User) error {
	return nil
}

func (ur DbUsersRepository) Update(id uuid.UUID, user models.User) error {
	return nil
}

func (ur DbUsersRepository) Delete(id uuid.UUID) error {
	return nil
}

func (ur DbUsersRepository) Exists(id uuid.UUID) (bool, error) {
	return true, nil
}

package repositories

import (
	"global-birthday-server/models"
	"global-birthday-server/mongodb"
	"os"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DbBirthdaysRepository struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func NewBirthdaysRepository() *DbBirthdaysRepository {
	mongoClient, _ := mongodb.Connect()
	return &DbBirthdaysRepository{
		client:     mongoClient,
		collection: mongoClient.Database(os.Getenv("MONGODB_NAME")).Collection("birthdays"),
	}
}

func (br DbBirthdaysRepository) Get() (*[]models.Birthday, error) {
	return nil, nil
}

func (br DbBirthdaysRepository) GetById(id uuid.UUID) (*models.Birthday, error) {
	return nil, nil
}

func (br DbBirthdaysRepository) Create(birthday models.Birthday) error {
	return nil
}

func (br DbBirthdaysRepository) Update(id uuid.UUID, birthday models.Birthday) error {
	return nil
}

func (br DbBirthdaysRepository) Delete(id uuid.UUID) error {
	return nil
}

func (br DbBirthdaysRepository) Exists(id uuid.UUID) (bool, error) {
	return true, nil
}

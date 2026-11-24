package repositories

import (
	"global-birthday-server/models"
	"global-birthday-server/mongodb"
	"os"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DbEventsRepository struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func NewEventsRepository() *DbEventsRepository {
	mongoClient, _ := mongodb.Connect()
	return &DbEventsRepository{
		client:     mongoClient,
		collection: mongoClient.Database(os.Getenv("MONGODB_NAME")).Collection("events"),
	}
}

func (er DbEventsRepository) Get() (*[]models.Event, error) {
	return nil, nil
}

func (er DbEventsRepository) GetById(id uuid.UUID) (*models.Event, error) {
	return nil, nil
}

func (er DbEventsRepository) Create(event models.Event) error {
	return nil
}

func (er DbEventsRepository) Update(id uuid.UUID, event models.Event) error {
	return nil
}

func (er DbEventsRepository) Delete(id uuid.UUID) error {
	return nil
}

func (er DbEventsRepository) Exists(id uuid.UUID) (bool, error) {
	return true, nil
}

package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/joaomarcuslf/sucellus/definitions"
	errors "github.com/joaomarcuslf/sucellus/errors"
	"github.com/joaomarcuslf/sucellus/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceRepository struct {
	collection *mongo.Collection
}

func NewServiceRepository(connection definitions.DatabaseClient) *ServiceRepository {
	collection, err := connection.Collection("services")

	if err != nil {
		panic(
			errors.FormatError(
				"REPOSITORY_ERROR",
				"(ServiceRepository) Error connecting to database",
				err,
			),
		)
	}

	return &ServiceRepository{
		collection,
	}
}

func (r *ServiceRepository) Query(ctx context.Context, filter bson.M) ([]models.Service, error) {
	var aux []models.Service

	cur, err := r.collection.Find(ctx, filter)

	if err != nil {
		return []models.Service{}, errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Could not Find element",
			err,
		)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var service models.Service
		err := cur.Decode(&service)

		if err != nil {
			return []models.Service{}, errors.FormatError(
				"REPOSITORY_ERROR",
				"(ServiceRepository) Error during decode",
				err,
			)
		}

		aux = append(aux, service)
	}

	if err := cur.Err(); err != nil {
		return []models.Service{}, errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Cursor error",
			err,
		)
	}

	return aux, nil
}

func (r *ServiceRepository) Insert(ctx context.Context, model models.Service) error {
	_, err := r.collection.InsertOne(ctx, model)

	if err != nil {
		return errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Could not insert new value",
			err,
		)
	}

	return nil
}

func (r *ServiceRepository) Set(ctx context.Context, uid primitive.ObjectID, model models.Service) error {
	model.UpdatedDate.Time = time.Now()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": uid}, bson.M{"$set": model})

	if err != nil {
		return errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Update error",
			err,
		)
	}

	return nil
}

func (r *ServiceRepository) Get(ctx context.Context, id string) (models.Service, error) {
	var aux models.Service

	uid, _ := primitive.ObjectIDFromHex(id)

	err := r.collection.FindOne(ctx, bson.M{"_id": uid}).Decode(&aux)

	if err != nil {
		return models.Service{}, errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Could not Find item",
			err,
		)
	}

	return aux, nil
}

func (r *ServiceRepository) Create(ctx context.Context, body io.Reader) (models.Service, error) {
	var aux models.Service

	err := json.NewDecoder(body).Decode(&aux)

	if err != nil {
		return models.Service{}, errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Error during decode",
			err,
		)
	}

	err = r.Validate(ctx, aux)

	if err != nil {
		return models.Service{}, errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Validation error:",
			err,
		)
	}

	aux.CreatedDate.Time = time.Now()
	aux.UpdatedDate.Time = time.Now()

	aux.UName = fmt.Sprintf("%s-%s", aux.Name, aux.CreatedDate.Time.Format("20060102150405"))

	err = r.Insert(ctx, aux)

	if err != nil {
		return models.Service{}, errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Insertion error",
			err,
		)
	}

	err = r.collection.FindOne(ctx, bson.M{"u_name": aux.UName}).Decode(&aux)

	if err != nil {
		return models.Service{}, errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Insertion error",
			err,
		)
	}

	return aux, nil
}

func (r *ServiceRepository) Update(ctx context.Context, id string, body io.Reader) error {
	var aux models.Service

	uid, _ := primitive.ObjectIDFromHex(id)

	err := json.NewDecoder(body).Decode(&aux)

	if err != nil {
		return errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Error during decode",
			err,
		)
	}

	return r.Set(ctx, uid, aux)
}

func (r *ServiceRepository) Delete(ctx context.Context, id string) error {
	uid, _ := primitive.ObjectIDFromHex(id)

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": uid})

	if err != nil {
		return errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Deletion error",
			err,
		)
	}

	return nil
}

func (r *ServiceRepository) Validate(ctx context.Context, model models.Service) error {
	if model.Name == "" {
		return fmt.Errorf("FieldValidation: [Name] is required")
	}

	if model.Url == "" {
		return fmt.Errorf("FieldValidation: [Url] is required")
	}

	if model.Port == 0 {
		return fmt.Errorf("FieldValidation: [Port] is required")
	}

	if model.Language == "" {
		return fmt.Errorf("FieldValidation: [Language] is required")
	}

	if string(model.Port) == os.Getenv("PORT") {
		return fmt.Errorf("FieldValidation: [Port] is already in use")
	}

	if model.PoolingInterval <= 100 {
		return fmt.Errorf("FieldValidation: [PoolingInterval] must be bigger than 100")
	}

	services, _ := r.Query(ctx, bson.M{"port": model.Port})

	if len(services) > 0 && services[0].ID != model.ID {
		return fmt.Errorf("FieldValidation: [Port] is already in use")
	}

	return nil
}

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

func NewServiceRepository(connection definitions.DatabaseClient) definitions.Repository {
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

func (r *ServiceRepository) Query(ctx context.Context, filter bson.M) ([]interface{}, error) {
	var aux []interface{}

	cur, err := r.collection.Find(ctx, filter)

	if err != nil {
		return nil, errors.FormatError(
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
			return nil, errors.FormatError(
				"REPOSITORY_ERROR",
				"(ServiceRepository) Error during decode",
				err,
			)
		}

		aux = append(aux, service)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Cursor error",
			err,
		)
	}

	return aux, nil
}

func (r *ServiceRepository) Insert(ctx context.Context, model interface{}) error {
	_, err := r.collection.InsertOne(ctx, model.(models.Service))

	if err != nil {
		return errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Could not insert new value",
			err,
		)
	}

	return nil
}

func (r *ServiceRepository) Get(ctx context.Context, id string) (interface{}, error) {
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

func (r *ServiceRepository) Set(ctx context.Context, uid primitive.ObjectID, model interface{}) error {
	aux := model.(models.Service)

	aux.UpdatedDate.Time = time.Now()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": uid}, bson.M{"$set": aux})

	if err != nil {
		return errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Update error",
			err,
		)
	}

	return nil
}

func (r *ServiceRepository) Create(ctx context.Context, body io.Reader) (interface{}, error) {
	var aux models.Service

	err := json.NewDecoder(body).Decode(&aux)

	if err != nil {
		return nil, errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Error during decode",
			err,
		)
	}

	err = r.Validate(ctx, aux)

	if err != nil {
		return nil, err
	}

	aux.CreatedDate.Time = time.Now()
	aux.UpdatedDate.Time = time.Now()

	aux.UName = fmt.Sprintf("%s-%s", aux.Name, aux.CreatedDate.Time.Format("20060102150405"))

	err = r.Insert(ctx, aux)

	if err != nil {
		return nil, errors.FormatError(
			"REPOSITORY_ERROR",
			"(ServiceRepository) Insertion error",
			err,
		)
	}

	err = r.collection.FindOne(ctx, bson.M{"u_name": aux.UName}).Decode(&aux)

	if err != nil {
		return nil, errors.FormatError(
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

func (r *ServiceRepository) Validate(ctx context.Context, model interface{}) error {
	aux := model.(models.Service)

	if aux.Name == "" {
		return errors.FormatError(
			"VALIDATION_ERROR",
			"FieldName: [Name]",
			fmt.Errorf("FieldValidation: [Name] is required"),
		)
	}

	if aux.Url == "" {
		return errors.FormatError(
			"VALIDATION_ERROR",
			"FieldName: [Url]",
			fmt.Errorf("FieldValidation: [Url] is required"),
		)
	}

	if aux.Port == 0 {
		return errors.FormatError(
			"VALIDATION_ERROR",
			"FieldName: [Port]",
			fmt.Errorf("FieldValidation: [Port] is required"),
		)
	}

	if aux.Language == "" {
		return errors.FormatError(
			"VALIDATION_ERROR",
			"FieldName: [Language]",
			fmt.Errorf("FieldValidation: [Language] is required"),
		)
	}

	if string(aux.Port) == os.Getenv("PORT") {
		return errors.FormatError(
			"VALIDATION_ERROR",
			"FieldName: [Port]",
			fmt.Errorf("FieldValidation: [Port] is already in use"),
		)
	}

	if aux.PoolingInterval <= 100 {
		return errors.FormatError(
			"VALIDATION_ERROR",
			"FieldName: [PoolingInterval]",
			fmt.Errorf("FieldValidation: [PoolingInterval] must be bigger than 100"),
		)
	}

	services, _ := r.Query(ctx, bson.M{"port": aux.Port})

	if len(services) > 0 && services[0].(models.Service).ID != aux.ID {
		return errors.FormatError(
			"VALIDATION_ERROR",
			"FieldName: [PoolingInterval]",
			fmt.Errorf("FieldValidation: [Port] is already in use"),
		)
	}

	return nil
}

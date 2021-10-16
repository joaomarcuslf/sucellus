package models

import (
	definition "github.com/joaomarcuslf/sucellus/definitions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	ID              primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            string              `json:"name" bson:"name"`
	Url             string              `json:"url" bson:"url"`
	Port            int                 `json:"port" bson:"port"`
	Language        string              `json:"language" bson:"language"`
	PoolingInterval int                 `json:"pooling_interval" bson:"pooling_interval"`
	Status          string              `json:"status" bson:"status"`
	EnvVars         []map[string]string `json:"env_vars" bson:"env_vars"`
	CreatedDate     definition.Datetime `json:"created_date" bson:"created_date"`
	UpdatedDate     definition.Datetime `json:"updated_date" bson:"updated_date"`
}

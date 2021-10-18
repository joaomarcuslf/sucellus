package models

import (
	definition "github.com/joaomarcuslf/sucellus/definitions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	ID              primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            string              `json:"name,omitempty" bson:"name,omitempty"`
	Url             string              `json:"url,omitempty" bson:"url,omitempty"`
	Port            int                 `json:"port,omitempty" bson:"port,omitempty"`
	Language        string              `json:"language,omitempty" bson:"language,omitempty"`
	PoolingInterval int                 `json:"pooling_interval,omitempty" bson:"pooling_interval,omitempty"`
	Status          string              `json:"status,omitempty" bson:"status,omitempty"`
	EnvVars         map[string]string   `json:"env_vars,omitempty" bson:"env_vars,omitempty"`
	CreatedDate     definition.Datetime `json:"created_date,omitempty" bson:"created_date,omitempty"`
	UpdatedDate     definition.Datetime `json:"updated_date,omitempty" bson:"updated_date,omitempty"`
	LastHead        string              `json:"last_head,omitempty" bson:"last_head,omitempty"`
	UName           string              `json:"u_name,omit" bson:"u_name,omitempty"`
}

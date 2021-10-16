package models

import (
	definition "github.com/joaomarcuslf/sucellus/definitions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Migration struct {
	ID   primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string              `json:"name" bson:"name"`
	Date definition.Datetime `json:"date" bson:"date"`
}

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credit struct {
	ID           primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`
	QuestionerID primitive.ObjectID `json:"questionerId"        bson:"questionerId"`
	Name         string             `json:"name" bson:"name"`
	Instagram    string             `json:"instagram"            bson:"instagram"`

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}

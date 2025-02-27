package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AttributeEmbedded struct {
	ID   *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name *string             `json:"name"          bson:"name"`
	Key  *int                `json:"key"           bson:"key"`
}

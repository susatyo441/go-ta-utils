package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductPhoto struct {
	ID        primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`
	Key       int                `json:"key"            bson:"key"`
	Photo     string             `json:"photo"            bson:"photo"`
	ProductID primitive.ObjectID `json:"productId"        bson:"productId"`

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}

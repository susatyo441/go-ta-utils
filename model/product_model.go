package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID         primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`
	Name       string             `json:"name"            bson:"name"`
	Category   AttributeEmbedded  `json:"category"            bson:"category"`
	CoverPhoto string             `json:"coverPhoto" bson:"coverPhoto"`
	Stock      int                `json:"stock" bson:"stock"`
	StoreID    primitive.ObjectID `json:"storeId"        bson:"storeId"`

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}

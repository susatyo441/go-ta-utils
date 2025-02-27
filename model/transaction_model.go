package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID         primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`
	Product    AttributeEmbedded  `json:"product"            bson:"product"`
	Category   AttributeEmbedded  `json:"category"            bson:"category"`
	Quantity   int                `json:"quantity"          bson:"quantity"`
	Price      int                `json:"price"          bson:"price"`
	TotalPrice int                `json:"totalPrice"          bson:"totalPrice"`

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}

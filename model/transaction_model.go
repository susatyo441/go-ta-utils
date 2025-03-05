package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionProductAttribute struct {
	ID       primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`
	Name     string             `json:"name"            bson:"name"`
	Price    int                `json:"price"          bson:"price"`
	Category AttributeEmbedded  `json:"category"            bson:"category"`
	Quantity int                `json:"quantity"          bson:"quantity"`
}

type Transaction struct {
	ID         primitive.ObjectID            `json:"_id,omitempty"        bson:"_id,omitempty"`
	Products   []TransactionProductAttribute `json:"products"            bson:"products"`
	TotalPrice int                           `json:"totalPrice"          bson:"totalPrice"`
	StoreID    primitive.ObjectID            `json:"storeId"        bson:"storeId"`

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductVariantsAttr struct {
	Name         string `json:"name" bson:"name"`
	Price        int    `json:"price" bson:"price"`
	CapitalPrice *int   `json:"capitalPrice" bson:"capitalPrice"`
	Stock        int    `json:"stock" bson:"stock"`
}

type Product struct {
	ID           primitive.ObjectID    `json:"_id,omitempty"        bson:"_id,omitempty"`
	Name         string                `json:"name"            bson:"name"`
	Category     AttributeEmbedded     `json:"category"            bson:"category"`
	CoverPhoto   string                `json:"coverPhoto" bson:"coverPhoto"`
	Stock        *int                  `json:"stock" bson:"stock"`
	Price        *int                  `json:"price" bson:"price"`
	StoreID      primitive.ObjectID    `json:"storeId"        bson:"storeId"`
	CapitalPrice *int                  `json:"capitalPrice" bson:"capitalPrice"`
	Variants     []ProductVariantsAttr `json:"variants" bson:"variants"`

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}

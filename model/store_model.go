package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	ID            primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`
	Name          string             `json:"name"            bson:"name"`
	Address       *string            `json:"address"            bson:"address"`
	LogoSmall     *string            `json:"logoSmall"  bson:"logoSmall"`
	LogoMedium    *string            `json:"logoMedium" bson:"logoMedium"`
	LogoBig       *string            `json:"logoBig"    bson:"logoBig"`
	PhotoModel    *string            `json:"photoModel" bson:"photoModel"`
	ForecastModel *string            `json:"forecastModel" bson:"forecastModel"`

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}

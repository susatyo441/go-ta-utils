package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                   primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`
	Name                 string             `json:"name"            bson:"name"`
	Store                AttributeEmbedded  `json:"store"            bson:"store"`
	PhoneNumber          *string            `json:"phoneNumber"          bson:"phoneNumber"`
	Email                string             `json:"email"                bson:"email"`
	Password             *string            `json:"password"             bson:"password"`
	ProfilePictureSmall  *string            `json:"profilePictureSmall"  bson:"profilePictureSmall"`
	ProfilePictureMedium *string            `json:"profilePictureMedium" bson:"profilePictureMedium"`
	ProfilePictureBig    *string            `json:"profilePictureBig"    bson:"profilePictureBig"`

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Changelog struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Action       string             `json:"action"        bson:"action"`
	Field        *string            `json:"field"         bson:"field"`
	OldValue     interface{}        `json:"oldValue"      bson:"oldValue"`
	NewValue     interface{}        `json:"newValue"      bson:"newValue"`
	ModifiedBy   string             `json:"modifiedBy"    bson:"modifiedBy"`
	ModifiedById primitive.ObjectID `json:"modifiedById"  bson:"modifiedById"`
	Object       string             `json:"object"        bson:"object"`
	ObjectId     primitive.ObjectID `json:"objectId"      bson:"objectId"`
	ObjectName   string             `json:"objectName"    bson:"objectName"`
	CreatedAt    time.Time          `json:"createdAt"     bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"     bson:"updatedAt"`
}

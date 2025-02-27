package parser

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

// Parse a struct into a map.
func StructToMap[T any](v T) (bson.M, error) {
	value := reflect.ValueOf(v)
	if !value.IsValid() {
		return bson.M{}, nil
	}

	jsonData, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err = bson.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}

	return result, nil
}

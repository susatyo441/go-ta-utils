package parser

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QueryTransform string

const (
	TransformString        QueryTransform = "string"
	TransformInt           QueryTransform = "int"
	TransformFloat32       QueryTransform = "float32"
	TransformFloat64       QueryTransform = "float64"
	TransformBool          QueryTransform = "bool"
	TransformArray         QueryTransform = "array"
	TransformObjectID      QueryTransform = "objectId"
	TransformObjectIDArray QueryTransform = "objectIdArray"
)

var queryTransformMap = map[string]QueryTransform{
	"string":        TransformString,
	"int":           TransformInt,
	"float32":       TransformFloat32,
	"float64":       TransformFloat64,
	"bool":          TransformBool,
	"array":         TransformArray,
	"objectId":      TransformObjectID,
	"objectIdArray": TransformObjectIDArray,
}

func getQueryTransform(transform string) (QueryTransform, error) {
	transformType, exist := queryTransformMap[transform]
	if !exist {
		return "", fmt.Errorf("invalid query transform: %s", transform)
	}
	return transformType, nil
}

func transformQueryValue(value string, transform string) interface{} {
	transformType, err := getQueryTransform(transform)
	if err != nil {
		fmt.Println("Error getting query transform:", err)
		return value
	}

	switch transformType {
	case TransformInt:
		var v, _ = strconv.Atoi(value)
		return v
	case TransformFloat32:
		var v, _ = strconv.ParseFloat(value, 32)
		return v
	case TransformFloat64:
		var v, _ = strconv.ParseFloat(value, 64)
		return v
	case TransformBool:
		var v, err = strconv.ParseBool(value)
		if err != nil {
			return nil //if value isnt string "false", (example: nil) it should return nil (case: optionally filtering boolean)
		}
		return v
	case TransformArray:
		var v []interface{}
		json.Unmarshal([]byte(value), &v)
		return v
	case TransformObjectID:
		var v, _ = primitive.ObjectIDFromHex(value)
		return v
	case TransformObjectIDArray:
		var v []interface{}
		var objectIdArr []primitive.ObjectID

		json.Unmarshal([]byte(value), &v)

		for _, val := range v {
			objectId, err := primitive.ObjectIDFromHex(val.(string))
			if err != nil {
				return v
			}

			objectIdArr = append(objectIdArr, objectId)
		}

		return objectIdArr
	default:
		return value
	}
}

func ParseQuery[T any](query map[string]string) (*T, error) {
	t := reflect.TypeFor[T]()

	parsedQuery := map[string]interface{}{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		key := field.Tag.Get("json")
		transform := field.Tag.Get("transform")
		parsedQuery[key] = transformQueryValue(query[key], transform)
	}

	jsonData, err := json.Marshal(parsedQuery)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil, err
	}

	var parsedData T
	err = json.Unmarshal(jsonData, &parsedData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	return &parsedData, nil
}

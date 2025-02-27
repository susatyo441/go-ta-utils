package pipeline

import (
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateRangeQuery(filter []int) bson.M {
	if len(filter) == 2 && filter[0] >= 0 && filter[1] >= 0 {
		return bson.M{
			"$gte": filter[0],
			"$lte": filter[1],
		}
	}

	if len(filter) == 2 && filter[1] >= 0 {
		return bson.M{
			"$lte": filter[1],
		}
	}

	if len(filter) == 1 && filter[0] >= 0 {
		return bson.M{
			"$gte": filter[0],
		}
	}

	return bson.M{
		"$ne": "vdjfPyv7ijO5vQeLIZmQHuzPO",
	}
}

func GenerateRangeQueryTwoVar(filter1 int, filter2 int) bson.M {
	if filter1 != 0 && filter2 != 0 {
		return bson.M{
			"$gte": filter1,
			"$lte": filter2,
		}
	}

	if filter2 != 0 {
		return bson.M{
			"$lte": filter2,
		}
	}

	if filter1 != 0 {
		return bson.M{
			"$gte": filter1,
		}
	}

	return bson.M{
		"$ne": "vdjfPyv7ijO5vQeLIZmQHuzPO",
	}
}

func GenerateExactFilter[T comparable](condition bool, filter T) primitive.M {
	if condition {
		return bson.M{"$eq": filter}
	} else {
		return bson.M{"$ne": "vdjfPyv7ijO5vQeLIZmQHuzPO"} //random string so its always true
	}
}

func GenerateObjectIdFilter(filter string) primitive.M {
	if filter == "" {
		return bson.M{"$ne": "vdjfPyv7ijO5vQeLIZmQHuzPO"} //random string so its always true
	}

	objId, err := primitive.ObjectIDFromHex(filter)
	if err != nil {
		return bson.M{"$eq": "djfPyv7ijO5vQeLIZmQHuzPO"}
	} else {
		return bson.M{"$eq": objId}
	}
}

func GenerateArrayFilter[T comparable](filter []T) primitive.M {
	var varType = reflect.TypeOf(filter).Kind()

	//check if variable type is slice AND the variable have length.
	//providing an empty slice will result to an error. (probably because it resolves as a zero value <nil>)
	if varType == reflect.Slice && len(filter) > 0 {
		return bson.M{"$in": filter}
	} else {
		return bson.M{"$ne": "vdjfPyv7ijO5vQeLIZmQHuzPO"} //random string so its always true
	}
}

func GenerateSearchCondition(search string) primitive.M {
	if search != "" {
		return bson.M{"$regex": search, "$options": "i"}
	} else {
		return bson.M{"$ne": false}
	}
}

func GenerateDateFilter(requestDate []int) primitive.M {
	if len(requestDate) < 1 {
		return bson.M{"$ne": "g#*a,Jvb&U.Bg"}
	}

	//No longer adds 24 hour minus one second to end. It should be added in FE
	if len(requestDate) == 2 && requestDate[0] >= 0 && requestDate[1] >= 0 {
		start := time.UnixMilli(int64(requestDate[0]))
		end := time.UnixMilli(int64(requestDate[1]))

		return bson.M{
			"$gte": start,
			"$lte": end,
		}
	}

	//If one date passed, the second argument will be 24 hour minus one second (same day) after the date
	if len(requestDate) == 1 && requestDate[0] >= 0 {
		start := time.UnixMilli(int64(requestDate[0]))
		start2 := time.UnixMilli(int64(requestDate[0]))

		start2 = start2.Add(24*time.Hour - time.Second)

		return bson.M{
			"$gte": start,
			"$lt":  start2,
		}
	}

	return bson.M{"$ne": "g#*a,Jvb&U.Bg"}
}

// Deprecated: use GenerateFacetOption func instead, so you no longer project the option with "$option._id"
func GenerateOptionFacet(
	fieldName string,
	optionsQuery bool,
	labelKey string,
	valueKey string,
) mongo.Pipeline {
	lk := labelKey
	if !strings.HasPrefix(lk, "$") {
		lk = "$" + lk
	}

	vk := valueKey
	if !strings.HasPrefix(vk, "$") {
		vk = "$" + vk
	}

	return NewPipelineBuilder().
		Match(bson.M{
			fieldName: bson.M{"$ne": nil},
			"$expr":   bson.M{"$eq": bson.A{optionsQuery, true}}}).
		Group(bson.D{
			{Key: "_id", Value: bson.M{"label": lk, "value": vk}},
		}).
		Sort(bson.M{"_id.label": 1}).
		Build()
}

// Generate option facet
func GenerateFacetOption(
	fieldName string,
	optionsQuery bool,
	labelKey string,
	valueKey string,
) mongo.Pipeline {
	lk := labelKey
	if !strings.HasPrefix(lk, "$") {
		lk = "$" + lk
	}

	vk := valueKey
	if !strings.HasPrefix(vk, "$") {
		vk = "$" + vk
	}

	return NewPipelineBuilder().
		Match(bson.M{
			fieldName: bson.M{"$ne": nil},
			"$expr":   bson.M{"$eq": bson.A{optionsQuery, true}}}).
		Group(bson.D{
			{Key: "_id", Value: bson.M{"label": lk, "value": vk}},
		}).
		Sort(bson.M{"_id.label": 1}).
		ReplaceRoot(bson.M{"newRoot": "$_id"}).
		Build()
}

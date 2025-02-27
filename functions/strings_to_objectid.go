package functions

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StringsToObjectId(ids []string) ([]primitive.ObjectID, int) {
	oids := make([]primitive.ObjectID, 0, len(ids))
	errors := 0

	for _, id := range ids {
		oid, castErr := primitive.ObjectIDFromHex(id)

		if castErr != nil {
			fmt.Println("Error converting", id, "to ObjectId:", castErr)
			errors += 1
		} else {
			oids = append(oids, oid)
		}
	}

	return oids, errors
}

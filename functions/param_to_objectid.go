package functions

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/susatyo441/go-ta-utils/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Helper function to convert a param to an ObjectID.
// Returns the ObjectID and an entity.HttpError if the conversion fails. Intended for use in controllers.
//
// EXAMPLE:
//
//		casted, castErr := ParamToObjectID(ctx, "myParamContainingId")
//	    // OR
//		casted, castErr := ParamToObjectID(ctx, "myParamContainingId", "Invalid ID")
//		if castErr != nil {
//		   return castErr.SendResponse(ctx)
//		}
func ParamToObjectID(
	ctx *fiber.Ctx,
	paramName string,
	errMsg ...string,
) (primitive.ObjectID, *entity.HttpError) {
	param := ctx.Params(paramName)

	castedParam, castErr := primitive.ObjectIDFromHex(param)
	if castErr == primitive.ErrInvalidHex {
		fmt.Println(castErr)
		if len(errMsg) == 0 {
			return castedParam, entity.BadRequest(paramName + " is not a valid ObjectId")
		} else {
			return castedParam, entity.BadRequest(errMsg[0])
		}
	} else if castErr != nil {
		fmt.Println(castErr)
		return castedParam, entity.InternalServerError("Error converting " + paramName + " to ObjectId")
	}

	return castedParam, nil
}

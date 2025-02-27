package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	internalvalidator "github.com/kittipat1413/go-common/framework/validator"
	"github.com/susatyo441/go-ta-utils/entity"
	"github.com/susatyo441/go-ta-utils/parser"
	customvalidator "github.com/susatyo441/go-ta-utils/validator/custom_validator"
)

// EXAMPLE USAGE:
// body, err := validation.ParseAndValidateBody[dto.MyDTO](ctx)
// OR
// query, err := validation.ParseAndValidateQuery[dto.MyQueryDTO](ctx)

// Send response in controller
// if err != nil {
// 	return err.SendResponse(ctx)
// }

func RawValidate[T any](data T) *entity.HttpError {
	// Create a new validator instance
	v, validatorErr := internalvalidator.NewValidator(
		// Use JSON tag name
		internalvalidator.WithTagNameFunc(internalvalidator.JSONTagNameFunc),

		// Register custom validator
		internalvalidator.WithCustomValidator(new(customvalidator.NotBlankValidator)),
	)

	// Return if there is an error when creating the validator
	if validatorErr != nil {
		return entity.InternalServerError(validatorErr.Error())
	}

	// Validate the data and return if there is an error
	if validateErr := v.ValidateStruct(data); validateErr != nil {
		return entity.BadRequest(validateErr.Error())
	}

	return nil
}

// Compacted common validation to reduce repeated code.
// Intended to be used in controller
func ParseAndValidateBody[T any](ctx *fiber.Ctx) (*T, *entity.HttpError) {
	// Parse the body
	var body T
	if err := ctx.BodyParser(&body); err != nil {
		return nil, entity.InternalServerError(err.Error())
	}

	// Validate the body and return error if there is any
	validationErr := RawValidate(body)
	if validationErr != nil {
		return nil, validationErr
	}

	// Return the parsed and validated body
	return &body, validationErr

}

// Compacted common validation to reduce repeated code.
// Intended to be used in controller
func ParseAndValidateQuery[T any](ctx *fiber.Ctx) (*T, *entity.HttpError) {
	// Parse the query
	rawQuery := ctx.Queries()

	// Transform query using "transform" tag
	// LINK: acts-utils/parser/query_parser.go
	parsedQuery, err := parser.ParseQuery[T](rawQuery)

	// Return if there is an error when parsing the query
	if err != nil {
		return nil, entity.BadRequest(err.Error())
	}

	// Validate the query and return if there is an error
	validationErr := RawValidate(parsedQuery)
	if validationErr != nil {
		return nil, validationErr
	}

	// Return the parsed and validated query
	return parsedQuery, nil
}

func ValidateStruct(obj any) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(obj)
}

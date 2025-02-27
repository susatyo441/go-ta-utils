package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Middleware to validate JWT
func ValidateJWT() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var jwtKey = []byte(os.Getenv("JWT_KEY"))
		tokenStr := ctx.Get("Authorization", "NOT_FOUND")
		if tokenStr == "NOT_FOUND" {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}

		// Extract the actual token by removing "Bearer " prefix
		tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)

		// Parse and validate the JWT
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm used
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			userId, _ := primitive.ObjectIDFromHex(claims["id"].(string))
			session, _ := primitive.ObjectIDFromHex(claims["session"].(string))
			store, _ := primitive.ObjectIDFromHex(claims["store"].(string))

			ctx.Locals(UserKey, userId)
			ctx.Locals(SessionKey, session)
			ctx.Locals(StoreKey, store)

		} else {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}

		return ctx.Next()
	}
}

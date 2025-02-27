package middleware

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
)

func SaveSingleImageMiddleware(folderName string, attribute string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// Get the uploaded file, FormFile attribute
		fileHeader, err := ctx.FormFile(attribute)
		if err != nil {
			return ctx.Next()
		}

		file, err := fileHeader.Open()
		if err != nil {
			return ctx.Next()
		}

		defer file.Close()

		// Decode the image
		img, _, err := image.Decode(file)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Error decoding image: "+err.Error())

		}

		absPath, _ := filepath.Abs(fmt.Sprintf("../acts-files/%s", folderName))
		// Save the resized image to disk
		outputFilePath := filepath.Join(
			absPath,
			fmt.Sprintf("%s.jpg", strconv.FormatInt(time.Now().Unix(), 10)),
		)
		relativePath := fmt.Sprintf(
			"/%s/%s",
			folderName,
			fmt.Sprintf("%s.jpg", strconv.FormatInt(time.Now().Unix(), 10)),
		)

		// Create the directory if it doesn't exist
		if err := os.MkdirAll(absPath, os.ModePerm); err != nil {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				"Error creating directory: "+err.Error(),
			)
		}

		saveErr := imaging.Save(img, outputFilePath)
		if saveErr != nil {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				"Error saving resized image: "+saveErr.Error(),
			)
		}

		key := ContextKey(attribute)

		ctx.Locals(key, relativePath)
		return ctx.Next()
	}
}

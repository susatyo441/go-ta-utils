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

// Deprecated: use middleware.SaveSingleImageMiddleware (for single file) since admin console doesnt need splitted image
func SplitImageMiddleware(folderName string, attribute string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// Get the uploaded file, FormFile attribute
		fileHeader, err := ctx.FormFile(attribute)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Error retrieving the file: "+err.Error())
		}

		file, err := fileHeader.Open()
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Error retrieving the file: "+err.Error())
		}

		defer file.Close()

		// Decode the image
		img, _, err := image.Decode(file)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Error decoding image: "+err.Error())

		}

		// Define sizes for resizing the image
		sizes := []struct {
			name          string
			width, height int
		}{
			{"small", 300, 300},
			{"medium", 600, 600},
			{"large", 1200, 1200},
		}

		// Process and save resized images
		for _, size := range sizes {
			// Resize the image to the specified dimensions
			resizedImg := imaging.Resize(img, size.width, size.height, imaging.Lanczos)

			companyCode := ctx.Locals(ContextKey("companyCode")).(string)

			absPath, _ := filepath.Abs(fmt.Sprintf("../acts-files/%s/%s", companyCode, folderName))
			// Save the resized image to disk
			outputFilePath := filepath.Join(
				absPath,
				fmt.Sprintf("%s.jpg", strconv.FormatInt(time.Now().Unix(), 10)+"-"+size.name),
			)

			// Create the directory if it doesn't exist
			if err := os.MkdirAll(absPath, os.ModePerm); err != nil {
				return fiber.NewError(
					fiber.StatusInternalServerError,
					"Error creating directory: "+err.Error(),
				)
			}

			err := imaging.Save(resizedImg, outputFilePath)
			if err != nil {
				return fiber.NewError(
					fiber.StatusInternalServerError,
					"Error saving resized image: "+err.Error(),
				)
			}

			key := ContextKey(size.name)

			ctx.Locals(key, outputFilePath)
		}

		return ctx.Next()
	}
}

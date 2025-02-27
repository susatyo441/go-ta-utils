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

func SaveMultiImageMiddleware(folderName string, attributes []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		for _, attribute := range attributes {
			// Get the uploaded file, FormFile attribute
			fileHeader, err := ctx.FormFile(attribute)
			if err != nil {
				continue
			}

			file, err := fileHeader.Open()
			if err != nil {
				continue
			}
			defer file.Close()

			// Decode the image
			img, _, err := image.Decode(file)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Error decoding image: "+err.Error())
			}

			absPath, _ := filepath.Abs(fmt.Sprintf("../acts-files/%s", folderName))
			// Gunakan UnixNano untuk menghasilkan nama file yang lebih unik
			timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
			outputFilePath := filepath.Join(absPath, fmt.Sprintf("%s.jpg", timestamp))
			relativePath := fmt.Sprintf("/%s/%s.jpg", folderName, timestamp)

			// Create the directory if it doesn't exist
			if err := os.MkdirAll(absPath, os.ModePerm); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Error creating directory: "+err.Error())
			}

			if err := imaging.Save(img, outputFilePath); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Error saving resized image: "+err.Error())
			}

			key := ContextKey(attribute)
			ctx.Locals(key, relativePath)
		}
		return ctx.Next()
	}
}

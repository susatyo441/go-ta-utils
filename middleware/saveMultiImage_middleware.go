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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveMultiImageMiddleware(folderName string, attributes []string, size PhotoSizeKey) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		storeId := ctx.Locals(StoreKey).(primitive.ObjectID)

		// Tentukan faktor pengurangan berdasarkan ukuran yang dipilih
		var scaleFactor float64
		switch size {
		case Low:
			scaleFactor = 0.25 // 75% dikurangi
		case Medium:
			scaleFactor = 0.50 // 50% dikurangi
		case High:
			scaleFactor = 1.0 // Tidak dikurangi
		default:
			scaleFactor = 1.0 // Default ke high jika tidak valid
		}

		for _, attribute := range attributes {
			// Get the uploaded file
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

			// Resize gambar jika perlu
			if scaleFactor < 1.0 {
				newWidth := int(float64(img.Bounds().Dx()) * scaleFactor)
				newHeight := int(float64(img.Bounds().Dy()) * scaleFactor)
				img = imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)
			}

			absPath, _ := filepath.Abs(fmt.Sprintf("../acts-files/%s/%s", storeId.Hex(), folderName))
			timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
			outputFilePath := filepath.Join(absPath, fmt.Sprintf("%s.jpg", timestamp))
			relativePath := fmt.Sprintf("/%s/%s.jpg", folderName, timestamp)

			// Create directory if it doesn't exist
			if err := os.MkdirAll(absPath, os.ModePerm); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Error creating directory: "+err.Error())
			}

			// Save the resized image
			if err := imaging.Save(img, outputFilePath); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Error saving resized image: "+err.Error())
			}

			key := ContextKey(attribute)
			ctx.Locals(key, relativePath)
		}
		return ctx.Next()
	}
}

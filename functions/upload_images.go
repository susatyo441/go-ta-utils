package functions

import (
	"errors"
	"fmt"
	"image"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/disintegration/imaging"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PhotoSizeKey defines size option
type PhotoSizeKey string

const (
	Low    PhotoSizeKey = "low"
	Medium PhotoSizeKey = "medium"
	High   PhotoSizeKey = "high"
)

func SaveMultiImages(storeId primitive.ObjectID, folderName string, attributes []string, files map[string]*multipart.FileHeader, size PhotoSizeKey) (map[string]string, error) {
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

	result := make(map[string]string)

	for _, attribute := range attributes {
		fileHeader, exists := files[attribute]
		if !exists || fileHeader == nil {
			continue // Skip kalau file-nya gak dikirim
		}

		file, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			return nil, fmt.Errorf("error decoding image: %w", err)
		}

		if scaleFactor < 1.0 {
			newWidth := int(float64(img.Bounds().Dx()) * scaleFactor)
			newHeight := int(float64(img.Bounds().Dy()) * scaleFactor)
			img = imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)
		}

		absPath, _ := filepath.Abs(fmt.Sprintf("../acts-files/%s/%s", storeId.Hex(), folderName))
		timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
		outputFilePath := filepath.Join(absPath, fmt.Sprintf("%s.jpg", timestamp))
		relativePath := fmt.Sprintf("/%s/%s.jpg", folderName, timestamp)

		if err := os.MkdirAll(absPath, os.ModePerm); err != nil {
			return nil, fmt.Errorf("error creating directory: %w", err)
		}

		if err := imaging.Save(img, outputFilePath); err != nil {
			return nil, fmt.Errorf("error saving resized image: %w", err)
		}

		result[attribute] = relativePath
	}

	if len(result) == 0 {
		return nil, errors.New("no images saved")
	}

	return result, nil
}

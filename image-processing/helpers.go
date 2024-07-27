package imageprocessing

import (
	"fmt"
	"image"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
)

func GetImageFiles(dirPath string) ([]string, error) {
	var imageFiles []string

	if err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("GetImageFiles: filepath.Walk %v", err)
		}

		if !info.IsDir() {
			imageFiles = append(imageFiles, path)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return imageFiles, nil
}

func LoadImages(imageFiles []string) ([]image.Image, error) {
	images := make([]image.Image, len(imageFiles))

	for idx, imagePath := range imageFiles {
		img, err := LoadImage(imagePath)
		if err != nil {
			return nil, err
		}

		images[idx] = img
	}

	return images, nil
}

func LoadImage_PNG(imagePath string) (image.Image, error) {
	imageFile, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("LoadImage: os.Open %v", err)
	}
	defer imageFile.Close()

	image, err := png.Decode(imageFile)
	if err != nil {
		return nil, fmt.Errorf("LoadImage: png.Decode %v", err)
	}

	return image, nil
}

func LoadImage(imagePath string) (image.Image, error) {
	imageFile, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("LoadImage: os.Open %v", err)
	}
	defer imageFile.Close()

	image, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, fmt.Errorf("LoadImage: image.Decode %v %v", imagePath, err)
	}

	return image, nil
}

package upload

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
)

// ImageConfig конфигурация обработки изображений
type ImageConfig struct {
	MaxWidth  uint
	MaxHeight uint
	Quality   int
}

// ResizeImage изменяет размер изображения
func ResizeImage(sourcePath, destPath string, maxWidth, maxHeight uint) error {
	// Открытие исходного файла
	file, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Декодирование изображения
	img, format, err := image.Decode(file)
	if err != nil {
		return err
	}

	// Изменение размера
	resized := resize.Thumbnail(maxWidth, maxHeight, img, resize.Lanczos3)

	// Создание файла назначения
	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Кодирование в зависимости от формата
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(out, resized, &jpeg.Options{Quality: 90})
	case "png":
		err = png.Encode(out, resized)
	default:
		return fmt.Errorf("unsupported image format: %s", format)
	}

	if err != nil {
		return err
	}

	logrus.Infof("Image resized: %s -> %s", sourcePath, destPath)
	return nil
}

// CreateThumbnail создает миниатюру изображения
func CreateThumbnail(sourcePath string, width, height uint) (string, error) {
	ext := filepath.Ext(sourcePath)
	base := strings.TrimSuffix(sourcePath, ext)
	thumbPath := fmt.Sprintf("%s_thumb%s", base, ext)

	err := ResizeImage(sourcePath, thumbPath, width, height)
	if err != nil {
		return "", err
	}

	return thumbPath, nil
}

// ValidateImage проверяет, является ли файл изображением
func ValidateImage(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, _, err = image.DecodeConfig(file)
	if err != nil {
		return fmt.Errorf("invalid image file")
	}

	return nil
}

// GetImageDimensions возвращает размеры изображения
func GetImageDimensions(path string) (int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}

	return config.Width, config.Height, nil
}

package upload

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Config конфигурация загрузки файлов
type Config struct {
	UploadDir      string
	MaxFileSize    int64
	AllowedTypes   []string
	GenerateUnique bool
}

var (
	config *Config
)

// Init инициализирует конфигурацию загрузки файлов
func Init() {
	uploadDir := os.Getenv("UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = "./uploads"
	}

	maxSize := int64(10 * 1024 * 1024) // 10MB default
	if size := os.Getenv("MAX_FILE_SIZE"); size != "" {
		// Parse size
	}

	config = &Config{
		UploadDir:      uploadDir,
		MaxFileSize:    maxSize,
		AllowedTypes:   []string{".jpg", ".jpeg", ".png", ".gif", ".pdf", ".doc", ".docx", ".txt"},
		GenerateUnique: true,
	}

	// Создание директории для загрузок
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		logrus.Errorf("Failed to create upload directory: %v", err)
	}

	logrus.Info("File upload service initialized")
}

// UploadFile загружает файл
func UploadFile(file *multipart.FileHeader) (*FileInfo, error) {
	if config == nil {
		Init()
	}

	// Проверка размера
	if file.Size > config.MaxFileSize {
		return nil, fmt.Errorf("file size exceeds maximum allowed size")
	}

	// Проверка типа файла
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !isAllowedType(ext) {
		return nil, fmt.Errorf("file type not allowed: %s", ext)
	}

	// Открытие файла
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Генерация имени файла
	filename := file.Filename
	if config.GenerateUnique {
		filename = generateUniqueFilename(file.Filename)
	}

	// Путь для сохранения
	destPath := filepath.Join(config.UploadDir, filename)

	// Создание файла
	dst, err := os.Create(destPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Копирование содержимого
	size, err := io.Copy(dst, src)
	if err != nil {
		return nil, err
	}

	// Вычисление MD5
	src.Seek(0, 0)
	hash := md5.New()
	io.Copy(hash, src)
	md5sum := hex.EncodeToString(hash.Sum(nil))

	fileInfo := &FileInfo{
		Filename:     filename,
		OriginalName: file.Filename,
		Path:         destPath,
		Size:         size,
		Extension:    ext,
		MimeType:     file.Header.Get("Content-Type"),
		MD5:          md5sum,
		UploadedAt:   time.Now(),
	}

	logrus.Infof("File uploaded: %s (%d bytes)", filename, size)
	return fileInfo, nil
}

// UploadMultiple загружает несколько файлов
func UploadMultiple(files []*multipart.FileHeader) ([]*FileInfo, error) {
	var results []*FileInfo
	var errors []string

	for _, file := range files {
		info, err := UploadFile(file)
		if err != nil {
			errors = append(errors, fmt.Sprintf("%s: %v", file.Filename, err))
			continue
		}
		results = append(results, info)
	}

	if len(errors) > 0 {
		return results, fmt.Errorf("some files failed: %s", strings.Join(errors, "; "))
	}

	return results, nil
}

// DeleteFile удаляет файл
func DeleteFile(filename string) error {
	path := filepath.Join(config.UploadDir, filename)
	
	if err := os.Remove(path); err != nil {
		return err
	}

	logrus.Infof("File deleted: %s", filename)
	return nil
}

// GetFilePath возвращает полный путь к файлу
func GetFilePath(filename string) string {
	return filepath.Join(config.UploadDir, filename)
}

// FileExists проверяет существование файла
func FileExists(filename string) bool {
	path := filepath.Join(config.UploadDir, filename)
	_, err := os.Stat(path)
	return err == nil
}

// isAllowedType проверяет, разрешен ли тип файла
func isAllowedType(ext string) bool {
	for _, allowed := range config.AllowedTypes {
		if ext == allowed {
			return true
		}
	}
	return false
}

// generateUniqueFilename генерирует уникальное имя файла
func generateUniqueFilename(original string) string {
	ext := filepath.Ext(original)
	name := strings.TrimSuffix(original, ext)
	
	// Очистка имени файла
	name = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			return r
		}
		return '_'
	}, name)

	// Добавление UUID
	id := uuid.New().String()[:8]
	timestamp := time.Now().Unix()
	
	return fmt.Sprintf("%s_%d_%s%s", name, timestamp, id, ext)
}

// FileInfo информация о загруженном файле
type FileInfo struct {
	Filename     string    `json:"filename"`
	OriginalName string    `json:"original_name"`
	Path         string    `json:"path"`
	Size         int64     `json:"size"`
	Extension    string    `json:"extension"`
	MimeType     string    `json:"mime_type"`
	MD5          string    `json:"md5"`
	UploadedAt   time.Time `json:"uploaded_at"`
}

// SetAllowedTypes устанавливает разрешенные типы файлов
func SetAllowedTypes(types []string) {
	if config == nil {
		Init()
	}
	config.AllowedTypes = types
}

// SetMaxFileSize устанавливает максимальный размер файла
func SetMaxFileSize(size int64) {
	if config == nil {
		Init()
	}
	config.MaxFileSize = size
}

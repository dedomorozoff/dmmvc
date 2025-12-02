package controllers

import (
	"dmmvc/internal/i18n"
	"dmmvc/internal/upload"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UploadSingleFile загружает один файл
// @Summary Загрузить файл
// @Description Загружает один файл на сервер
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} APIResponse{data=upload.FileInfo}
// @Failure 400 {object} APIResponse
// @Router /api/upload/file [post]
// @Security SessionAuth
func UploadSingleFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "No file provided",
		})
		return
	}

	fileInfo, err := upload.UploadFile(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "File uploaded successfully",
		Data:    fileInfo,
	})
}

// UploadMultipleFiles загружает несколько файлов
// @Summary Загрузить несколько файлов
// @Description Загружает несколько файлов на сервер
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param files formData file true "Files to upload" multiple
// @Success 200 {object} APIResponse{data=[]upload.FileInfo}
// @Failure 400 {object} APIResponse
// @Router /api/upload/files [post]
// @Security SessionAuth
func UploadMultipleFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Failed to parse form",
		})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "No files provided",
		})
		return
	}

	fileInfos, err := upload.UploadMultiple(files)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
			Data:    fileInfos, // Partial success
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Files uploaded successfully",
		Data:    fileInfos,
	})
}

// UploadImage загружает изображение и создает миниатюру
// @Summary Загрузить изображение
// @Description Загружает изображение и создает миниатюру
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image to upload"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Router /api/upload/image [post]
// @Security SessionAuth
func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "No image provided",
		})
		return
	}

	// Загрузка файла
	fileInfo, err := upload.UploadFile(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Проверка, что это изображение
	if err := upload.ValidateImage(fileInfo.Path); err != nil {
		upload.DeleteFile(fileInfo.Filename)
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Invalid image file",
		})
		return
	}

	// Создание миниатюры
	thumbPath, err := upload.CreateThumbnail(fileInfo.Path, 300, 300)
	if err != nil {
		c.JSON(http.StatusOK, APIResponse{
			Success: true,
			Message: "Image uploaded but thumbnail creation failed",
			Data:    fileInfo,
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Image uploaded successfully",
		Data: gin.H{
			"file":      fileInfo,
			"thumbnail": filepath.Base(thumbPath),
		},
	})
}

// DeleteUploadedFile удаляет загруженный файл
// @Summary Удалить файл
// @Description Удаляет загруженный файл
// @Tags upload
// @Accept json
// @Produce json
// @Param filename path string true "Filename"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Router /api/upload/file/{filename} [delete]
// @Security SessionAuth
func DeleteUploadedFile(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Filename is required",
		})
		return
	}

	if !upload.FileExists(filename) {
		c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Error:   "File not found",
		})
		return
	}

	if err := upload.DeleteFile(filename); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to delete file",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "File deleted successfully",
	})
}

// DownloadFile скачивает файл
// @Summary Скачать файл
// @Description Скачивает загруженный файл
// @Tags upload
// @Produce application/octet-stream
// @Param filename path string true "Filename"
// @Success 200 {file} binary
// @Failure 404 {object} APIResponse
// @Router /api/upload/file/{filename} [get]
// @Security SessionAuth
func DownloadFile(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Filename is required",
		})
		return
	}

	if !upload.FileExists(filename) {
		c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Error:   "File not found",
		})
		return
	}

	filePath := upload.GetFilePath(filename)
	c.File(filePath)
}

// UploadPage отображает страницу загрузки файлов
func UploadPage(c *gin.Context) {
	locale := i18n.GetLocale(c)
	c.HTML(http.StatusOK, "pages/upload.html", gin.H{
		"title":  i18nT(c, "upload.title"),
		"locale": i18nLocale(c),
		"T":      i18n.TFunc(locale),
	})
}

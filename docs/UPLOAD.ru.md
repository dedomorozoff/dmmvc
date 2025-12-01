[English](UPLOAD.md) | **Русский**

# Загрузка файлов

DMMVC включает встроенную поддержку загрузки файлов с возможностями обработки изображений.

## Возможности

- **Одиночная и множественная загрузка** - Загрузка одного или нескольких файлов
- **Валидация типов файлов** - Ограничение разрешенных типов
- **Лимиты размера** - Контроль максимального размера файла
- **Уникальные имена** - Автоматическая генерация уникальных имен
- **Обработка изображений** - Изменение размера и создание миниатюр
- **MD5 контрольные суммы** - Проверка целостности файлов

## Быстрый старт

### 1. Настройка загрузки

Отредактируйте файл `.env`:

```env
# Директория загрузки
UPLOAD_DIR=./uploads

# Максимальный размер файла (10MB)
MAX_FILE_SIZE=10485760
```

### 2. Запуск сервера

```bash
make run
```

### 3. Доступ к странице загрузки

Откройте: **http://localhost:8080/upload**

## Использование

### Загрузка одного файла

```go
import "dmmvc/internal/upload"

func UploadHandler(c *gin.Context) {
    file, _ := c.FormFile("file")
    
    fileInfo, err := upload.UploadFile(file)
    if err != nil {
        // Обработка ошибки
    }
    
    // Файл успешно загружен
    // fileInfo содержит: filename, size, path и т.д.
}
```

### Загрузка нескольких файлов

```go
form, _ := c.MultipartForm()
files := form.File["files"]

fileInfos, err := upload.UploadMultiple(files)
```

## Обработка изображений

### Изменение размера

```go
import "dmmvc/internal/upload"

// Изменить размер до макс 800x600
err := upload.ResizeImage(
    "uploads/image.jpg",
    "uploads/image_resized.jpg",
    800,
    600,
)
```

### Создание миниатюры

```go
// Создать миниатюру 300x300
thumbPath, err := upload.CreateThumbnail(
    "uploads/image.jpg",
    300,
    300,
)
```

## API эндпоинты

- `POST /api/upload/file` - Загрузить один файл
- `POST /api/upload/files` - Загрузить несколько файлов
- `POST /api/upload/image` - Загрузить изображение с миниатюрой
- `GET /api/upload/file/:filename` - Скачать файл
- `DELETE /api/upload/file/:filename` - Удалить файл

### Пример запроса

```bash
# Загрузить один файл
curl -X POST http://localhost:8080/api/upload/file \
  -F "file=@document.pdf"

# Загрузить несколько файлов
curl -X POST http://localhost:8080/api/upload/files \
  -F "files=@file1.jpg" \
  -F "files=@file2.jpg"
```

## Пример контроллера

```go
func UploadAvatar(c *gin.Context) {
    file, _ := c.FormFile("avatar")
    
    // Загрузить файл
    fileInfo, err := upload.UploadFile(file)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // Проверить, что это изображение
    if err := upload.ValidateImage(fileInfo.Path); err != nil {
        upload.DeleteFile(fileInfo.Filename)
        c.JSON(400, gin.H{"error": "Неверное изображение"})
        return
    }
    
    // Создать миниатюру
    thumbPath, _ := upload.CreateThumbnail(fileInfo.Path, 150, 150)
    
    // Обновить аватар пользователя в БД
    user.Avatar = fileInfo.Filename
    database.DB.Save(&user)
    
    c.JSON(200, gin.H{"message": "Аватар загружен"})
}
```

## Управление файлами

### Удаление файла

```go
err := upload.DeleteFile("filename.jpg")
```

### Проверка существования

```go
if upload.FileExists("filename.jpg") {
    // Файл существует
}
```

## Разрешенные типы файлов

По умолчанию:
- Изображения: `.jpg`, `.jpeg`, `.png`, `.gif`
- Документы: `.pdf`, `.doc`, `.docx`, `.txt`

### Пользовательские типы

```go
upload.SetAllowedTypes([]string{
    ".jpg", ".png", ".gif",  // Изображения
    ".pdf", ".doc", ".docx", // Документы
    ".zip", ".rar",          // Архивы
})
```

## Лимиты размера

По умолчанию: 10MB

```go
// Лимит 5MB
upload.SetMaxFileSize(5 * 1024 * 1024)
```

## Безопасность

### Валидация типов файлов

Всегда проверяйте типы файлов:

```go
ext := filepath.Ext(file.Filename)
if !isAllowedType(ext) {
    return errors.New("тип файла не разрешен")
}
```

### Санитизация имен

Имена файлов автоматически очищаются:
- Удаляются специальные символы
- Добавляется уникальный ID
- Включается временная метка

Пример: `document_1701234567_a1b2c3d4.pdf`

## Лучшие практики

1. **Валидируйте типы файлов** - Проверяйте расширение и содержимое
2. **Устанавливайте лимиты размера** - Предотвращайте большие загрузки
3. **Используйте уникальные имена** - Избегайте конфликтов
4. **Сохраняйте метаданные** - Храните информацию о файлах в БД
5. **Очищайте старые файлы** - Реализуйте очистку файлов
6. **Сканируйте на вирусы** - Используйте антивирус в production
7. **Используйте CDN** - Раздавайте файлы через CDN

## HTML форма

```html
<form action="/api/upload/file" method="post" enctype="multipart/form-data">
    <input type="file" name="file" required>
    <button type="submit">Загрузить</button>
</form>
```

## JavaScript загрузка

```javascript
async function uploadFile(file) {
    const formData = new FormData();
    formData.append('file', file);
    
    const response = await fetch('/api/upload/file', {
        method: 'POST',
        body: formData
    });
    
    const result = await response.json();
    if (result.success) {
        console.log('Файл загружен:', result.data.filename);
    }
}
```

## Production конфигурация

```env
# Используйте абсолютный путь
UPLOAD_DIR=/var/www/uploads

# Больший лимит размера
MAX_FILE_SIZE=52428800  # 50MB
```

## Решение проблем

### Загрузка не работает

1. Проверьте, что директория загрузки существует и доступна для записи
2. Проверьте, что размер файла в пределах лимита
3. Проверьте, что тип файла разрешен

### Обработка изображений не работает

1. Убедитесь, что изображение валидно
2. Проверьте, что формат поддерживается (JPEG, PNG)
3. Проверьте наличие свободного места на диске

## Ресурсы

- [Gin File Upload](https://gin-gonic.com/docs/examples/upload-file/)
- [Go Image Package](https://pkg.go.dev/image)
- [nfnt/resize](https://github.com/nfnt/resize)

## Следующие шаги

- Реализовать задачу очистки файлов
- Добавить сканирование на вирусы
- Интегрировать с облачным хранилищем
- Добавить оптимизацию изображений
- Реализовать версионирование файлов

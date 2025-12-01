**English** | [Русский](UPLOAD.ru.md)

# File Upload

DMMVC includes built-in support for file uploads with image processing capabilities.

## Features

- **Single & multiple file uploads** - Upload one or many files
- **File type validation** - Restrict allowed file types
- **Size limits** - Control maximum file size
- **Unique filenames** - Automatic unique name generation
- **Image processing** - Resize and create thumbnails
- **MD5 checksums** - File integrity verification

## Quick Start

### 1. Configure Upload Settings

Edit `.env` file:

```env
# Upload directory
UPLOAD_DIR=./uploads

# Maximum file size (10MB)
MAX_FILE_SIZE=10485760
```

### 2. Start Server

```bash
make run
```

### 3. Access Upload Page

Open: **http://localhost:8080/upload**

## Usage

### Upload Single File

```go
import "dmmvc/internal/upload"

func UploadHandler(c *gin.Context) {
    file, _ := c.FormFile("file")
    
    fileInfo, err := upload.UploadFile(file)
    if err != nil {
        // Handle error
    }
    
    // File uploaded successfully
    // fileInfo contains: filename, size, path, etc.
}
```

### Upload Multiple Files

```go
form, _ := c.MultipartForm()
files := form.File["files"]

fileInfos, err := upload.UploadMultiple(files)
```

### Upload with Custom Configuration

```go
// Set allowed file types
upload.SetAllowedTypes([]string{".jpg", ".png", ".pdf"})

// Set max file size (5MB)
upload.SetMaxFileSize(5 * 1024 * 1024)

// Upload file
fileInfo, err := upload.UploadFile(file)
```

## Image Processing

### Resize Image

```go
import "dmmvc/internal/upload"

// Resize to max 800x600
err := upload.ResizeImage(
    "uploads/image.jpg",
    "uploads/image_resized.jpg",
    800,
    600,
)
```

### Create Thumbnail

```go
// Create 300x300 thumbnail
thumbPath, err := upload.CreateThumbnail(
    "uploads/image.jpg",
    300,
    300,
)
```

### Validate Image

```go
err := upload.ValidateImage("uploads/file.jpg")
if err != nil {
    // Not a valid image
}
```

### Get Image Dimensions

```go
width, height, err := upload.GetImageDimensions("uploads/image.jpg")
```

## API Endpoints

DMMVC includes example upload endpoints:

- `POST /api/upload/file` - Upload single file
- `POST /api/upload/files` - Upload multiple files
- `POST /api/upload/image` - Upload image with thumbnail
- `GET /api/upload/file/:filename` - Download file
- `DELETE /api/upload/file/:filename` - Delete file

### Example Request

```bash
# Upload single file
curl -X POST http://localhost:8080/api/upload/file \
  -F "file=@document.pdf"

# Upload multiple files
curl -X POST http://localhost:8080/api/upload/files \
  -F "files=@file1.jpg" \
  -F "files=@file2.jpg"

# Upload image
curl -X POST http://localhost:8080/api/upload/image \
  -F "image=@photo.jpg"
```

## Controller Example

```go
func UploadAvatar(c *gin.Context) {
    file, _ := c.FormFile("avatar")
    
    // Upload file
    fileInfo, err := upload.UploadFile(file)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // Validate it's an image
    if err := upload.ValidateImage(fileInfo.Path); err != nil {
        upload.DeleteFile(fileInfo.Filename)
        c.JSON(400, gin.H{"error": "Invalid image"})
        return
    }
    
    // Create thumbnail
    thumbPath, _ := upload.CreateThumbnail(fileInfo.Path, 150, 150)
    
    // Update user avatar in database
    user.Avatar = fileInfo.Filename
    user.AvatarThumb = filepath.Base(thumbPath)
    database.DB.Save(&user)
    
    c.JSON(200, gin.H{
        "message": "Avatar uploaded",
        "avatar": fileInfo.Filename,
    })
}
```

## File Management

### Delete File

```go
err := upload.DeleteFile("filename.jpg")
```

### Check if File Exists

```go
if upload.FileExists("filename.jpg") {
    // File exists
}
```

### Get File Path

```go
path := upload.GetFilePath("filename.jpg")
// Returns: ./uploads/filename.jpg
```

## Allowed File Types

Default allowed types:
- Images: `.jpg`, `.jpeg`, `.png`, `.gif`
- Documents: `.pdf`, `.doc`, `.docx`, `.txt`

### Custom Allowed Types

```go
upload.SetAllowedTypes([]string{
    ".jpg", ".png", ".gif",  // Images
    ".pdf", ".doc", ".docx", // Documents
    ".zip", ".rar",          // Archives
    ".mp4", ".avi",          // Videos
})
```

## File Size Limits

Default: 10MB

### Custom Size Limit

```go
// 5MB limit
upload.SetMaxFileSize(5 * 1024 * 1024)

// 50MB limit
upload.SetMaxFileSize(50 * 1024 * 1024)
```

## Security

### File Type Validation

Always validate file types:

```go
ext := filepath.Ext(file.Filename)
if !isAllowedType(ext) {
    return errors.New("file type not allowed")
}
```

### Filename Sanitization

Filenames are automatically sanitized:
- Special characters removed
- Unique ID added
- Timestamp included

Example: `document_1701234567_a1b2c3d4.pdf`

### Content Type Validation

For images, validate actual content:

```go
if err := upload.ValidateImage(path); err != nil {
    // Not a real image
}
```

## Best Practices

1. **Validate file types** - Check both extension and content
2. **Set size limits** - Prevent large file uploads
3. **Use unique filenames** - Avoid conflicts
4. **Store metadata** - Save file info in database
5. **Clean up old files** - Implement file cleanup
6. **Scan for viruses** - Use antivirus in production
7. **Use CDN** - Serve files from CDN in production

## HTML Form Example

```html
<form action="/api/upload/file" method="post" enctype="multipart/form-data">
    <input type="file" name="file" required>
    <button type="submit">Upload</button>
</form>

<!-- Multiple files -->
<form action="/api/upload/files" method="post" enctype="multipart/form-data">
    <input type="file" name="files" multiple required>
    <button type="submit">Upload Files</button>
</form>
```

## JavaScript Upload Example

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
        console.log('File uploaded:', result.data.filename);
    }
}
```

## Progress Tracking

```javascript
function uploadWithProgress(file) {
    const formData = new FormData();
    formData.append('file', file);
    
    const xhr = new XMLHttpRequest();
    
    xhr.upload.addEventListener('progress', (e) => {
        if (e.lengthComputable) {
            const percent = (e.loaded / e.total) * 100;
            console.log(`Upload progress: ${percent}%`);
        }
    });
    
    xhr.open('POST', '/api/upload/file');
    xhr.send(formData);
}
```

## Database Integration

Store file metadata in database:

```go
type Upload struct {
    gorm.Model
    Filename     string
    OriginalName string
    Size         int64
    MimeType     string
    MD5          string
    UserID       uint
}

func SaveUpload(fileInfo *upload.FileInfo, userID uint) {
    upload := Upload{
        Filename:     fileInfo.Filename,
        OriginalName: fileInfo.OriginalName,
        Size:         fileInfo.Size,
        MimeType:     fileInfo.MimeType,
        MD5:          fileInfo.MD5,
        UserID:       userID,
    }
    database.DB.Create(&upload)
}
```

## Production Configuration

```env
# Use absolute path
UPLOAD_DIR=/var/www/uploads

# Larger size limit
MAX_FILE_SIZE=52428800  # 50MB

# Use separate storage
# UPLOAD_DIR=/mnt/storage/uploads
```

## Cloud Storage Integration

For production, consider using cloud storage:

### AWS S3

```go
import "github.com/aws/aws-sdk-go/service/s3"

func UploadToS3(file *multipart.FileHeader) error {
    // Upload to S3
}
```

### Google Cloud Storage

```go
import "cloud.google.com/go/storage"

func UploadToGCS(file *multipart.FileHeader) error {
    // Upload to GCS
}
```

## Troubleshooting

### Upload fails

1. Check upload directory exists and is writable
2. Verify file size is within limit
3. Check file type is allowed

### Image processing fails

1. Ensure image is valid
2. Check image format is supported (JPEG, PNG)
3. Verify sufficient disk space

### Permission denied

```bash
# Make upload directory writable
chmod 755 uploads
```

## Resources

- [Gin File Upload](https://gin-gonic.com/docs/examples/upload-file/)
- [Go Image Package](https://pkg.go.dev/image)
- [nfnt/resize](https://github.com/nfnt/resize)

## Next Steps

- Implement file cleanup job
- Add virus scanning
- Integrate with cloud storage
- Add image optimization
- Implement file versioning

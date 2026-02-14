package file

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// 1 MB = 1024 * 1024 = 2^20 bytes
// 10 MB = 10 * (2^20) bytes
const maxUploadSize = 10 << 20

var allowedTypes = map[string]bool{
	".jpg": true,
	".png": true,
	".pdf": true,
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check file size
	if file.Size > maxUploadSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file too large",
		})
		return
	}

	// sanitize the filename, helps prevent path traversal attacks
	filename := filepath.Base(file.Filename) // strips dir path and only keep file name

	// check file type
	ext := strings.ToLower(filepath.Ext(filename))
	if !allowedTypes[ext] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file type not allowed",
		})
		return
	}

	dst := filepath.Join("uploads", filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "file uploaded",
		"filename": filename,
	})
}

func UploadFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	files := form.File["files"]
	var filenames []string

	for _, f := range files {
		filename := filepath.Base(f.Filename)
		dst := filepath.Join("uploads", filename)
		if err := c.SaveUploadedFile(f, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		filenames = append(filenames, filename)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "files uploaded",
		"filenames": filenames,
	})
}

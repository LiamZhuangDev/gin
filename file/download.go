package file

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) {
	filename := filepath.Base(c.Param("filename"))
	fullpath := filepath.Join("uploads", filename)

	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "File not found",
		})
		return
	}

	// Set the header
	c.Header("Content-Disposition", "attachment; filename="+filename) // don't display in browser, download it instead. Also use this name in Save dialog.

	// Send the file
	c.File(fullpath)
}

package file_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/LiamZhuangDev/gin/file"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestDownloadFile(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// prepare file
	os.Mkdir("uploads", 0755)
	defer os.RemoveAll("uploads")

	filename := "a.txt"
	content := "hello world"
	err := os.WriteFile(filepath.Join("uploads", filename), []byte(content), 0644)
	require.NoError(t, err)

	// router
	r := gin.Default()
	r.GET("/download/:filename", file.DownloadFile)

	// request
	req := httptest.NewRequest(http.MethodGet, "/download/"+filename, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assert
	require.Equal(t, http.StatusOK, w.Code)

	require.Equal(t,
		"attachment; filename="+filename,
		w.Header().Get("Content-Disposition"))

	body, err := io.ReadAll(w.Body)
	require.NoError(t, err)
	require.Equal(t, content, string(body))
}

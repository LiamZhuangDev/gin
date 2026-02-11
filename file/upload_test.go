package file_test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/LiamZhuangDev/gin/file"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestUploadFile(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// ensure upload dir
	os.Mkdir("uploads", 0755)
	defer os.RemoveAll("uploads")

	// router
	r := gin.Default()
	r.POST("/upload", file.UploadFile)

	// create multipart body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "test.txt")
	require.NoError(t, err)

	_, err = io.Copy(part, bytes.NewBufferString("hello world"))
	require.NoError(t, err)

	writer.Close()

	// create request
	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// record response
	w := httptest.NewRecorder()

	// serve
	r.ServeHTTP(w, req)

	// assertion
	require.Equal(t, http.StatusOK, w.Code)

	var res map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	require.Equal(t, "test.txt", res["filename"])

	// check file saved
	_, err = os.Stat(filepath.Join("uploads", "test.txt"))
	require.NoError(t, err)
}

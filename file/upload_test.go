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

func TestUploadFiles(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// ensure upload dir
	os.Mkdir("uploads", 0755)
	defer os.RemoveAll("uploads")

	// router
	r := gin.Default()
	r.POST("/upload", file.UploadFiles)

	// create multipart body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	p1, err := writer.CreateFormFile("files", "a.txt")
	require.NoError(t, err)

	_, err = io.Copy(p1, bytes.NewBufferString("aaa"))
	require.NoError(t, err)

	p2, err := writer.CreateFormFile("files", "b.txt")
	require.NoError(t, err)

	_, err = io.Copy(p2, bytes.NewBufferString("bbb"))
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

	var resp struct {
		Message    string   `json:"message"`
		Filesnames []string `json:"filenames"`
	}

	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	require.Equal(t, "files uploaded", resp.Message)
	require.ElementsMatch(t, []string{"a.txt", "b.txt"}, resp.Filesnames)

	for _, name := range []string{"a.txt", "b.txt"} {
		_, err := os.Stat(filepath.Join("uploads", name))
		require.NoError(t, err)
	}
}

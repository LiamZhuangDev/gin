package pb_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	pb "github.com/LiamZhuangDev/gin/user_proto/v1"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestCreateUserProto(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/users", createUserProto)

	reqBody := &pb.CreateUserRequest{
		Username: "alice",
		Email:    "alice@test.com",
		Age:      18,
	}
	data, err := proto.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/x-protobuf")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assert
	require.Equal(t, http.StatusOK, w.Code)

	var resp pb.CreateUserResponse
	err = proto.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.Equal(t, "alice", resp.User.Username)
	require.Equal(t, "alice@test.com", resp.User.Email)
	require.Equal(t, int32(18), resp.User.Age)
}

func createUserProto(c *gin.Context) {
	// 读取原始数据
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 反序列化
	var req pb.CreateUserRequest
	if err := proto.Unmarshal(data, &req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 处理请求
	user := &pb.User{
		Id:       1,
		Username: req.Username,
		Email:    req.Email,
		Age:      req.Age,
		Active:   true,
	}

	resp := &pb.CreateUserResponse{
		User:    user,
		Success: true,
		Message: "User created successfully",
	}

	// 返回 Protobuf 响应
	data, _ = proto.Marshal(resp)
	c.Data(200, "application/x-protobuf", data)
}

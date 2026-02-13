package grpc

import (
	"log"
	"testing"
	"time"

	user_rpcv1 "github.com/LiamZhuangDev/gin/user_rpc/v1"
	"github.com/stretchr/testify/require"
)

func TestUserServiceClient(t *testing.T) {
	// Start the gRPC server in a separate goroutine
	go StartGRPCServer()

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	// Create a new gRPC client
	client, err := NewUserServiceClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Test GetUser method
	_, err = client.GetUser(1)
	require.NoError(t, err, "GetUser should not return an error")

	// Test BatchCreateUsers method
	users := []*user_rpcv1.User{
		{Username: "Alice", Email: "alice@example.com", Age: 25},
		{Username: "Bob", Email: "bob@example.com", Age: 30},
	}
	successCount, err := client.BatchCreateUsers(users)
	require.NoError(t, err, "BatchCreateUsers should not return an error")
	require.Equal(t, int32(len(users)), successCount, "Success count should match number of users sent")
}

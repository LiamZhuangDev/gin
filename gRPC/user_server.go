package grpc

import (
	"context"
	"fmt"
	"io"
	"time"

	user_rpcv1 "github.com/LiamZhuangDev/gin/user_rpc/v1"
)

type UserServer struct {
	user_rpcv1.UnimplementedUserServiceServer
}

// GetUser implements a unary RPC method to retrieve user information.
func (s *UserServer) GetUser(ctx context.Context, req *user_rpcv1.GetUserRequest) (*user_rpcv1.GetUserResponse, error) {
	// TODO: query user from database
	user := &user_rpcv1.User{
		Id:       req.Id,
		Username: "John Doe",
		Email:    "john.doe@example.com",
		Age:      30,
		Active:   true,
	}

	return &user_rpcv1.GetUserResponse{
		User: user,
	}, nil
}

// StreamUsers implements a server streaming RPC method to stream user data.
func (s *UserServer) StreamUsers(req *user_rpcv1.StreamUsersRequest, stream user_rpcv1.UserService_StreamUsersServer) error {
	for i := range 5 {
		user := &user_rpcv1.User{
			Id:       int64(i),
			Username: "User " + fmt.Sprint(i),
			Email:    "user" + fmt.Sprint(i) + "@example.com",
			Age:      20 + int32(i),
			Active:   i%2 == 0,
		}
		if err := stream.Send(&user_rpcv1.StreamUsersResponse{User: user}); err != nil {
			return err
		}
	}
	return nil
}

// BatchCreateUsers implements a client streaming RPC method for batch user creation.
func (s *UserServer) BatchCreateUsers(stream user_rpcv1.UserService_BatchCreateUsersServer) error {
	var users []*user_rpcv1.User

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&user_rpcv1.BatchCreateUsersResponse{
				SuccessCount: int32(len(users)),
			})
		}
		if err != nil {
			return err
		}

		user := &user_rpcv1.User{
			Username: req.Username,
			Email:    req.Email,
			Age:      req.Age,
		}
		users = append(users, user)
	}
}

// Chat implements a bidirectional streaming RPC method for real-time chat.
func (s *UserServer) Chat(stream user_rpcv1.UserService_ChatServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		resp := &user_rpcv1.ChatResponse{
			UserId:    req.UserId,
			Message:   "echo: " + req.Message,
			Timestamp: time.Now().Unix(),
		}

		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}

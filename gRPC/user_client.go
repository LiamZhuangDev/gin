package grpc

import (
	"context"
	"fmt"

	user_rpcv1 "github.com/LiamZhuangDev/gin/user_rpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	conn   *grpc.ClientConn
	client user_rpcv1.UserServiceClient
}

func NewUserServiceClient(addr string) (*UserServiceClient, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}
	return &UserServiceClient{
		conn:   conn,
		client: user_rpcv1.NewUserServiceClient(conn),
	}, nil
}

func (c *UserServiceClient) Close() {
	c.conn.Close()
}

func (c *UserServiceClient) GetUser(id int64) (*user_rpcv1.User, error) {
	resp, err := c.client.GetUser(context.Background(), &user_rpcv1.GetUserRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("could not get user: %v", err)
	}
	return resp.User, nil
}

func (c *UserServiceClient) BatchCreateUsers(users []*user_rpcv1.User) (int32, error) {
	stream, err := c.client.BatchCreateUsers(context.Background())
	if err != nil {
		return 0, fmt.Errorf("could not start batch create: %v", err)
	}

	for _, user := range users {
		req := &user_rpcv1.BatchCreateUsersRequest{
			Username: user.Username,
			Email:    user.Email,
			Age:      user.Age,
		}
		if err := stream.Send(req); err != nil {
			return 0, fmt.Errorf("could not send user: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return 0, fmt.Errorf("could not receive response: %v", err)
	}
	return resp.SuccessCount, nil
}

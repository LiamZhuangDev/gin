package grpc

import (
	"log"
	"net"

	user_rpcv1 "github.com/LiamZhuangDev/gin/user_rpc/v1"
	"google.golang.org/grpc"
)

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	user_rpcv1.RegisterUserServiceServer(grpcServer, &UserServer{})

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

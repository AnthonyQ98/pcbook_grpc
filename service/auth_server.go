package service

import (
	"context"
	"pcbook/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	userStore UserStore
	pb.UnimplementedAuthServiceServer
	jwtManager *JWTManager
}

func NewAuthServer(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	return &AuthServer{userStore: userStore, jwtManager: jwtManager}
}

func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := server.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user %s not found", req.GetUsername())
	}
	if !user.CheckPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.Unauthenticated, "password is incorrect")
	}
	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate token: %v", err)
	}
	return &pb.LoginResponse{AccessToken: token}, nil
}

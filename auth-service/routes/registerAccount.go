package routes

import (
	"auth-service/auth"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthServer) RegisterAccount(ctx context.Context, req *auth.RegisterAccountRequest) (*emptypb.Empty, error) {
	//logic goes here
	return &emptypb.Empty{}, nil
}

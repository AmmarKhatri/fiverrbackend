package routes

import (
	"auth-service/auth"
	"auth-service/helpers"
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*emptypb.Empty, error) {
	token, err := helpers.GenerateToken()
	if err != nil {
		return nil, err
	}
	fmt.Println(token)
	//logic goes here
	return &emptypb.Empty{}, nil
}

package routes

import (
	"auth-service/auth"
	"auth-service/helpers"
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthServer) RegisterAccount(ctx context.Context, req *auth.RegisterAccountRequest) (*emptypb.Empty, error) {
	token, err := helpers.GenerateExternalToken()
	if err != nil {
		return nil, err
	}
	fmt.Println("External Token:", token)
	//logic goes here

	return &emptypb.Empty{}, nil
}

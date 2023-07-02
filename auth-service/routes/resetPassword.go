package routes

import (
	"auth-service/auth"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthServer) ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*emptypb.Empty, error) {
	//logic goes here
	return &emptypb.Empty{}, nil
}

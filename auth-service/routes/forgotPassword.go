package routes

import (
	"auth-service/auth"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthServer) ForgotPassword(ctx context.Context, req *auth.ForgotPasswordRequest) (*emptypb.Empty, error) {
	//logic goes here
	return &emptypb.Empty{}, nil
}

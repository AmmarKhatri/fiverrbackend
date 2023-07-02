package routes

import (
	"auth-service/auth"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthServer) Validate2FA(ctx context.Context, req *auth.Validate2FARequest) (*emptypb.Empty, error) {
	//logic goes here
	return &emptypb.Empty{}, nil
}

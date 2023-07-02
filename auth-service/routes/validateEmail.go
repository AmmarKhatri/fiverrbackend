package routes

import (
	"auth-service/auth"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthServer) ValidateEmail(ctx context.Context, req *auth.ValidateEmailRequest) (*emptypb.Empty, error) {
	//logic goes here
	return &emptypb.Empty{}, nil
}

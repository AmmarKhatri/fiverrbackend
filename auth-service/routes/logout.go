package routes

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthServer) Logout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	//logic goes here
	return &emptypb.Empty{}, nil
}

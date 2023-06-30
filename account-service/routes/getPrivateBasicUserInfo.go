package routes

import (
	"account-service/account"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

func (s *AccountServer) GetPrivateBasicUserInfo(ctx context.Context, req *empty.Empty) (*account.PrivateBasicUserInfo, error) {
	//logic goes here
	return &account.PrivateBasicUserInfo{}, nil
}

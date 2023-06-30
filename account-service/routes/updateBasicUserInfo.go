package routes

import (
	"account-service/account"
	"context"
)

func (s *AccountServer) UpdateBasicUserInfo(ctx context.Context, req *account.UpdatePrivateBasicUserInfoRequest) (*account.PrivateBasicUserInfo, error) {
	//logic goes here
	return &account.PrivateBasicUserInfo{}, nil
}

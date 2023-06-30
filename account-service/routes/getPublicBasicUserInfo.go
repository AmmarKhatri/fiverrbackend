package routes

import (
	"account-service/account"
	"context"
)

func (s *AccountServer) GetPublicBasicUserInfo(ctx context.Context, req *account.GetPublicBasicUserInfoRequest) (*account.PublicBasicUserInfo, error) {
	//logic goes here
	return &account.PublicBasicUserInfo{}, nil
}

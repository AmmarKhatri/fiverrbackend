package routes

import (
	"account-service/account"
	"context"
)

func (s *AccountServer) GetPrivateProfileSection(ctx context.Context, req *account.GetPrivateProfileSectionRequest) (*account.ProfileSectionResponse, error) {
	//logic goes here
	return &account.ProfileSectionResponse{}, nil
}

package routes

import (
	"account-service/account"
	"context"
)

func (s *AccountServer) UpdateProfileSection(ctx context.Context, req *account.UpdateProfileSectionRequest) (*account.ProfileSectionResponse, error) {
	//logic goes here
	return &account.ProfileSectionResponse{}, nil
}

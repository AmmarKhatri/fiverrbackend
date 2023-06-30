package routes

import (
	"account-service/account"
	"context"
)

func (s *AccountServer) GetPublicProfileSection(ctx context.Context, req *account.GetPublicProfileSectionRequest) (*account.ProfileSectionResponse, error) {
	//logic goes here
	return &account.ProfileSectionResponse{}, nil
}

package routes

import (
	"account-service/account"
	"context"
)

func (s *AccountServer) GetCustomerReviews(ctx context.Context, req *account.GetCustomerReviewsRequest) (*account.GetCustomerReviewsResponse, error) {
	//logic goes here
	return &account.GetCustomerReviewsResponse{}, nil
}

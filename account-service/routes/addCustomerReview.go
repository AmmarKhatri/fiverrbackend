package routes

import (
	"account-service/account"
	"context"
)

func (s *AccountServer) AddCustomerReview(ctx context.Context, req *account.AddCustomerReviewRequest) (*account.CustomerReview, error) {
	//logic goes here
	return &account.CustomerReview{}, nil
}

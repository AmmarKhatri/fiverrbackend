package routes

import (
	"context"
	"scheduler/schedule"
)

func (s *SchedulerServer) ListCustomerDays(ctx context.Context, req *schedule.ListDaysRequest) (*schedule.ListCustomerDaysResponse, error) {
	//logic goes here
	return &schedule.ListCustomerDaysResponse{}, nil
}

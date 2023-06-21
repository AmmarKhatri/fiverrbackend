package routes

import (
	"context"
	"fmt"
	"scheduler/schedule"
)

func (s *SchedulerServer) ListProviderDays(ctx context.Context, req *schedule.ListDaysRequest) (*schedule.ListProviderDaysResponse, error) {
	fmt.Println(&req)
	//logic goes here
	return &schedule.ListProviderDaysResponse{}, nil
}

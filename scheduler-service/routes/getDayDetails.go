package routes

import (
	"context"
	"scheduler/schedule"
)

func (s *SchedulerServer) GetDayDetails(ctx context.Context, req *schedule.DayInput) (*schedule.DayDetails, error) {
	//logic goes here
	return &schedule.DayDetails{}, nil
}

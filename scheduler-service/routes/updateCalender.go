package routes

import (
	"context"
	"scheduler/schedule"
)

func (s *SchedulerServer) UpdateCalendar(ctx context.Context, req *schedule.UpdateCalendarRequest) (*schedule.UpdateCalendarResponse, error) {
	//logic goes here
	return &schedule.UpdateCalendarResponse{
		Days: make([]*schedule.ProviderDay, 0),
	}, nil
}

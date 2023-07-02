package routes

import (
	"context"
	"scheduler/schedule"

	"google.golang.org/genproto/googleapis/type/date"
)

func (s *SchedulerServer) GetDayDetails(ctx context.Context, req *schedule.GetDayDetailsRequest) (*schedule.DayDetails, error) {
	//logic goes here

	return &schedule.DayDetails{
		//date needs a valid value when being output. Cannot be null
		CalendarDate: &date.Date{
			Year:  req.CalendarDate.Year,
			Month: req.CalendarDate.Month,
			Day:   req.CalendarDate.Day,
		},
	}, nil
}

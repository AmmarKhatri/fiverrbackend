package routes

import (
	"context"
	"scheduler/schedule"
)

func (s *SchedulerServer) AddAppointment(ctx context.Context, req *schedule.AddAppointmentRequest) (*schedule.Appointment, error) {
	//logic goes here
	return &schedule.Appointment{}, nil
}

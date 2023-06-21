package routes

import (
	"context"
	"scheduler/schedule"

	"github.com/golang/protobuf/ptypes/empty"
)

func (s *SchedulerServer) CancelAppointment(ctx context.Context, req *schedule.CancelAppointmentRequest) (*empty.Empty, error) {
	//logic goes here
	return &empty.Empty{}, nil
}

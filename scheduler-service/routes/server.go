package routes

import "scheduler/schedule"

type SchedulerServer struct {
	schedule.SchedulerServer
}

func NewSchedulerServer() *SchedulerServer {
	//define connections to your databases here
	return &SchedulerServer{}
}

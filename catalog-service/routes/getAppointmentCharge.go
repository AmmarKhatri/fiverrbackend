package routes

import (
	"catalog-service/catalog"
	"context"
)

func (s *CatalogServer) GetAppointmentCharge(ctx context.Context, req *catalog.GetAppointmentChargeRequest) (*catalog.AppointmentCharge, error) {
	//write logic here
	return &catalog.AppointmentCharge{}, nil
}

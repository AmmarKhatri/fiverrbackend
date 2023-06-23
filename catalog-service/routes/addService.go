package routes

import (
	"catalog-service/catalog"
	"context"
)

func (s *CatalogServer) AddService(ctx context.Context, req *catalog.AddServiceRequest) (*catalog.Service, error) {
	//write logic here
	return &catalog.Service{}, nil
}

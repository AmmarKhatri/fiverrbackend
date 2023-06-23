package routes

import (
	"catalog-service/catalog"
	"context"
)

func (s *CatalogServer) EditService(ctx context.Context, req *catalog.EditServiceRequest) (*catalog.Service, error) {
	//write logic here
	return &catalog.Service{}, nil
}

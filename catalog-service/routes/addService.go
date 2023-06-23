package routes

import (
	"catalog-service/catalog"
	"context"
	"fmt"
)

func (s *CatalogServer) AddService(ctx context.Context, req *catalog.AddServiceRequest) (*catalog.Service, error) {
	//write logic here
	fmt.Println(req.Price)
	return &catalog.Service{Price: req.Price}, nil
}

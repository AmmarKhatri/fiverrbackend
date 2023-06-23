package routes

import (
	"catalog-service/catalog"
	"context"
)

func (s *CatalogServer) GetCatalog(ctx context.Context, req *catalog.GetCatalogRequest) (*catalog.GetCatalogResponse, error) {
	//write logic here
	return &catalog.GetCatalogResponse{}, nil
}

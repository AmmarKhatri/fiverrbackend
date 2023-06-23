package routes

import (
	"catalog-service/catalog"
	"context"
)

func (s *CatalogServer) GetClientPrice(ctx context.Context, req *catalog.GetClientPriceRequest) (*catalog.ClientPrice, error) {
	//write logic here
	return &catalog.ClientPrice{}, nil
}

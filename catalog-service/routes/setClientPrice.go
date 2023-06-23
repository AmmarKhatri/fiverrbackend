package routes

import (
	"catalog-service/catalog"
	"context"
)

func (s *CatalogServer) SetClientPrice(ctx context.Context, req *catalog.SetClientPriceRequest) (*catalog.ClientPrice, error) {
	//write logic here
	return &catalog.ClientPrice{}, nil
}

package routes

import (
	"catalog-service/catalog"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

func (s *CatalogServer) DeleteService(ctx context.Context, req *catalog.DeleteServiceRequest) (*empty.Empty, error) {
	//write logic here
	return &empty.Empty{}, nil
}

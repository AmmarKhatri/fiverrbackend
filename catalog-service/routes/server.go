package routes

import "catalog-service/catalog"

type CatalogServer struct {
	catalog.CatalogServer
}

func NewCatalogServer() *CatalogServer {
	//define connections to your databases here
	return &CatalogServer{}
}

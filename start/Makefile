SCHEDULER_BINARY=schedulerApp
GRAPH_BINARY=graphApp
CATALOG_BINARY=catalogApp
ACCOUNT_BINARY=accountApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build: build_graph build_scheduler build_catalog build_account
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_graph: builds the graph binary as a linux executable
build_graph:
	@echo "Building graph binary..."
	cd ../graph-gateway && env GOOS=linux CGO_ENABLED=0 go build -o ${GRAPH_BINARY} .
	@echo "Done!"

## build_scheduler: builds the scheduler binary as a linux executable
build_scheduler:
	@echo "Building scheduler binary..."
	cd ../scheduler-service && env GOOS=linux CGO_ENABLED=0 go build -o ${SCHEDULER_BINARY} .
	@echo "Done!"

## build_catalog: builds the catalog binary as a linux executable
build_catalog:
	@echo "Building catalog binary..."
	cd ../catalog-service && env GOOS=linux CGO_ENABLED=0 go build -o ${CATALOG_BINARY}
	@echo "Done!"

## build_account: builds the account binary as a linux executable
build_account:
	@echo "Building account binary..."
	cd ../account-service && env GOOS=linux CGO_ENABLED=0 go build -o ${ACCOUNT_BINARY}
	@echo "Done!"
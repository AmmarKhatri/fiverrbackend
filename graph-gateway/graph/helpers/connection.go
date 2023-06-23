package helpers

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ScheduleConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("scheduler-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return conn, err
}

func CommsConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("comms-service:50002", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return conn, err
}

func CatalogConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("catalog-service:50003", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return conn, err
}

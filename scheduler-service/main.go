package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	routes "scheduler/routes"
	"scheduler/schedule"

	"google.golang.org/grpc"
)

func main() {
	//logging starts
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// //loading the .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	//intializing tcp connection
	port := os.Getenv("PORT")
	fmt.Println("Port: ", port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	//grpc server initiated
	s := grpc.NewServer()
	srv := routes.NewSchedulerServer()
	schedule.RegisterSchedulerServer(s, srv)
	log.Printf("gRPC Server started on port %s", port)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to listen for gRPC: %v", err)
		}
	}()

	//Wait for control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block until a signal is received
	<-ch

	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Stopping the listener")
	lis.Close()
	fmt.Println("End of Program")
}
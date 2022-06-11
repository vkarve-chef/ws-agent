package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "ws-agent/io.chef"

	"google.golang.org/grpc"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		os.Exit(0)
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 56001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChefInfraServer(grpcServer, &chefInfraServer{})
	log.Printf("GRPC server listening on %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

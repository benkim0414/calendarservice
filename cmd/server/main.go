package main

import (
	"context"
	"fmt"
	"log"
	"net"

	calendarpb "github.com/benkim0414/calendarservice/pb"
	"github.com/benkim0414/calendarservice/pkg/calendarservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":50051"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ctx := context.Background()
	svc, err := calendarservice.New(ctx)
	if err != nil {
		log.Fatalf("failed to initialize the service: %v", err)
	}

	s := grpc.NewServer()
	calendarpb.RegisterCalendarServer(s, svc)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

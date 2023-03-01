package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"service/auth/auth_service"
	g "service/auth/global"
	load "service/auth/load"
	"service/auth/service"
)

func main() {
	load.Info()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", g.CFG.CurrentMicroservice.IP, g.CFG.CurrentMicroservice.Port))
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	auth_service.RegisterAuthServer(s, service.New())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

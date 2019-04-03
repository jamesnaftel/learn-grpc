package main

import (
	"flag"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	port := flag.String("port", "3001", "Port to listen on")

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	if err != nil {
		log.Fatalf("error creating listen: %v", err)
	}

	server := grpc.NewServer()
	server.Serve(lis)

}

package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	pb "github.com/jamesnaftel/learn-grpc/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct{}

type podcast struct {
	Name   string
	Author string
	Length int32
}

type database map[string]podcast

func (s *server) GetPodcast(ctx context.Context, name *pb.PodcastRequest) (*pb.PodcastResponse, error) {

	db := initDatabase()
	p := db[name.GetName()]
	return &pb.PodcastResponse{Podcast: &pb.Podcast{Name: p.Name, Author: p.Author, Length: p.Length}}, nil
}

func (s *server) GetPodcasts(ctx context.Context, _ *pb.Empty) (*pb.PodcastsResponse, error) {

	db := initDatabase()

	p := []*pb.Podcast{}
	for _, val := range db {
		p = append(p, &pb.Podcast{Name: val.Name, Author: val.Author, Length: val.Length})
	}

	return &pb.PodcastsResponse{Podcasts: p}, nil
}

func main() {
	port := flag.String("port", "3001", "Port to listen on")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	if err != nil {
		log.Fatalf("error creating listen: %v", err)
	}

	gs := grpc.NewServer()
	pb.RegisterPodcastsServer(gs, &server{})
	gs.Serve(lis)

}

func initDatabase() database {
	p := database{
		"SE Daily: Graphql": podcast{"SE Daily: Graphql", "Jeff Meyerson", 30},
		"SE Daily: GRPC":    podcast{"SE Daily: GRPC", "Jeff Meyerson", 40},
		"Security Now #708": podcast{"Security Now #708", "SteveG", 20},
	}

	return p
}

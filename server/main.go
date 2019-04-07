package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	pb "github.com/jamesnaftel/learn-grpc/api"
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

func (s *server) GetPodcasts(_ *pb.Empty, stream pb.Podcasts_GetPodcastsServer) error {

	db := initDatabase()

	for _, val := range db {
		err := stream.Send(&pb.Podcast{Name: val.Name, Author: val.Author, Length: val.Length})
		if err != nil {
			return fmt.Errorf("error sending podcast: %v", err)
		}
	}

	return nil
}

func (s *server) AddPodcast(_ context.Context, p *pb.Podcast) (*pb.Podcast, error) {
	return p, nil
}

func main() {
	port := flag.String("port", "3001", "Port to listen on")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating lis: %v", err)
		os.Exit(1)
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

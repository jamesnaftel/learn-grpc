package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	proto "github.com/golang/protobuf/proto"
	pb "github.com/jamesnaftel/learn-grpc/api"
	"google.golang.org/grpc"
)

type server struct {
	databaseToServe string
}

type database []pb.Podcast

func (s *server) GetByName(ctx context.Context, name *pb.ByNameRequest) (*pb.Podcast, error) {

	db := initDatabase(s.databaseToServe)
	for _, val := range db {
		if name.GetName() == val.Name {
			return &val, nil
		}
	}

	return &pb.Podcast{}, nil
}

func (s *server) List(_ *pb.Empty, stream pb.Podcasts_ListServer) error {

	db := initDatabase(s.databaseToServe)

	for _, val := range db {
		err := stream.Send(&val)
		if err != nil {
			return fmt.Errorf("error sending podcast: %v", err)
		}
	}

	return nil
}

func (s *server) Add(_ context.Context, p *pb.Podcast) (*pb.Podcast, error) {

	b, err := proto.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal podcast: %v", err)
	}

	f, err := os.OpenFile(s.databaseToServe, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return nil, fmt.Errorf("unable to open database file: %v", err)
	}

	err = binary.Write(f, binary.LittleEndian, int64(len(b)))
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("failed to write item length: %v", err)
	}

	_, err = f.Write(b)
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("failed to write item: %v", err)
	}

	if err = f.Close(); err != nil {
		return nil, fmt.Errorf("failed to close db: %v", err)
	}

	return p, nil
}

func main() {
	port := flag.String("port", "3001", "Port to listen on")
	dbFileName := flag.String("dbfile", "podcast.db", "Name of database file to store and retrieve podcasts")
	flag.Parse()

	createDatabase(*dbFileName)

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating lis: %v", err)
		os.Exit(1)
	}

	gs := grpc.NewServer()
	pb.RegisterPodcastsServer(gs, &server{databaseToServe: *dbFileName})
	gs.Serve(lis)

}

func createDatabase(dbFileName string) {
	f, err := os.OpenFile(dbFileName, os.O_CREATE, 0664)
	if err != nil {
		fmt.Printf("failed to create podcast database: %v\n", err)
		return
	}
	if err = f.Close(); err != nil {
		fmt.Printf("createDatabase: failed to close database: %v", err)
	}
}

func initDatabase(dbFileName string) database {

	//read all contents of file
	b, err := ioutil.ReadFile(dbFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading %s: %v", dbFileName, err)
		os.Exit(1)
	}

	const sizeOfint64 int = 8

	pdb := database{}
	for {
		if len(b) == 0 {
			break
		}

		var messageLen int64
		err := binary.Read(bytes.NewReader(b[:sizeOfint64]), binary.LittleEndian, &messageLen)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading message size: %v", err)
			os.Exit(1)
		}

		b = b[sizeOfint64:]

		podcast := pb.Podcast{}
		err = proto.Unmarshal(b[:messageLen], &podcast)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error unmarshalling: %v\n", err)
			os.Exit(1)
		}

		pdb = append(pdb, podcast)

		b = b[messageLen:]
	}

	return pdb
}

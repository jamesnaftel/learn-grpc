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
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type database []pb.Podcast

type server struct {
	databaseToServe string
	db              database
}

func (s *server) GetByName(ctx context.Context, name *pb.ByNameRequest) (*pb.Podcast, error) {
	for _, val := range s.db {
		if name.GetName() == val.Name {
			return &val, nil
		}
	}

	return &pb.Podcast{}, nil
}

func (s *server) List(_ *pb.Empty, stream pb.Podcasts_ListServer) error {
	for _, val := range s.db {
		err := stream.Send(&val)
		if err != nil {
			log.Errorf("error in list function: %+v", err)
			return fmt.Errorf("error sending podcast: %v", err)
		}
	}

	return nil
}

func (s *server) Add(_ context.Context, p *pb.Podcast) (*pb.Podcast, error) {

	//Add to File
	b, err := proto.Marshal(p)
	if err != nil {
		log.Errorf("unable to marshal podcast: %+v", err)
		return nil, fmt.Errorf("add failed")
	}

	f, err := os.OpenFile(s.databaseToServe, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Errorf("unable to open database file: %v", err)
		return nil, fmt.Errorf("add failed")
	}

	err = binary.Write(f, binary.LittleEndian, int64(len(b)))
	if err != nil {
		f.Close()
		log.Errorf("failed to write item length: %v", err)
		return nil, fmt.Errorf("add failed")
	}

	_, err = f.Write(b)
	if err != nil {
		f.Close()
		log.Errorf("failed to write item: %v", err)
		return nil, fmt.Errorf("add failed")
	}

	if err = f.Close(); err != nil {
		log.Errorf("failed to close db: %v", err)
		return nil, fmt.Errorf("add failed")
	}

	//Add to in memory data structure
	s.db = append(s.db, *p)

	return p, nil
}

func main() {
	port := flag.String("port", "3001", "Port to listen on")
	dbFileName := flag.String("dbfile", "podcast.db", "Name of database file to store and retrieve podcasts")
	flag.Parse()

	server := server{databaseToServe: *dbFileName}

	err := createDatabase(server.databaseToServe)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}

	server.db, err = initDatabase(server.databaseToServe)
	if err != nil {
		log.Fatalf("failed to initalize database: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	if err != nil {
		log.Fatalf("error creating lis: %v", err)
	}

	gs := grpc.NewServer()
	pb.RegisterPodcastsServer(gs, &server)
	gs.Serve(lis)

}

func createDatabase(dbFileName string) error {
	f, err := os.OpenFile(dbFileName, os.O_CREATE, 0664)
	if err != nil {
		log.Fatalf("createDatabase: error creating database: %v", err)
		return err
	}
	if err = f.Close(); err != nil {
		log.Fatalf("createDatabase: error closing database: %v", err)
		return err
	}

	return nil
}

func initDatabase(dbFileName string) (database, error) {

	//read all contents of file
	b, err := ioutil.ReadFile(dbFileName)
	if err != nil {
		log.Fatalf("error reading %s: %v", dbFileName, err)
		return nil, err
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
			log.Fatalf("error reading message size: %v", err)
			return nil, err
		}

		b = b[sizeOfint64:]

		podcast := pb.Podcast{}
		err = proto.Unmarshal(b[:messageLen], &podcast)
		if err != nil {
			log.Fatalf("error unmarshalling: %v\n", err)
			return nil, err
		}

		pdb = append(pdb, podcast)

		b = b[messageLen:]
	}

	return pdb, nil
}

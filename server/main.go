package main

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	api "github.com/jamesnaftel/learn-grpc/api"
	log "github.com/sirupsen/logrus"
)

func main() {
	mypodcast := &api.Podcast{Name: "SecurityNow", Author: "Steve", Length: (2 << 20)}

	data, err := proto.Marshal(mypodcast)
	if err != nil {
		log.Errorf("error marshalling podcast %v", err)
	}

	newMyPodcast := &api.Podcast{}
	err = proto.Unmarshal(data, newMyPodcast)
	if err != nil {
		log.Errorf("error unmarshalling podcast %v", err)
	}

	fmt.Println(newMyPodcast)

}

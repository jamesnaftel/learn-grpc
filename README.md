# learn-grpc

## Go setup
go get -u google.golang.org/grpc

### Generate podcasts.pb.go file w/local tools
1. Install protoc
1. go get -u github.com/golang/protobuf/protoc-gen-go
1. cd api/ && protoc --go_out=plugins=grpc:. -I. podcasts.proto

### Generate podcasts.pb.go w/docker
docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --go_out=plugins=grpc:. -I. podcasts.proto

## Define the service
Create a .proto file that contains the service information
```
service TestService {
    ...
}
```



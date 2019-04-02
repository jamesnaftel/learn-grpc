gen:
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) znly/protoc --go_out=plugins=grpc:. -I. podcasts.proto

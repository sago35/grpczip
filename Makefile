# vi: ft=make
.PHONY: proto test get ci docker

proto:
	go get github.com/golang/protobuf/protoc-gen-go
	protoc -I . grpczip.proto --lile-server_out=. --go_out=plugins=grpc:${GOPATH}/src

test:
	go test -v ./... -cover

get:
	go get -u -t ./...

ci: get test

docker:
	docker build . -t grpczip:latest

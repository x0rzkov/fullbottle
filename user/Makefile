
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto

.PHONY: build
build: proto

	CGO_ENABLED=0 GOOS=linux go build -o user-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

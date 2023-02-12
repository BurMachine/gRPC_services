
.PHONY = gen
gen:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	protoc -I api/proto --go_out=pkg/api --go-grpc_out==grpc:pkg/api api/proto/first_service_pb.proto


run:
	go run cmd/main.go


.PHONY = gen
gen:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc -I api/proto --go_out=pkg/api --go-grpc_out==grpc:pkg/api api/proto/first_service_pb.proto


run:
	go run cmd/main.go

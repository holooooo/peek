grpc:
	protoc --go_out=src/gen --go-grpc_out=src/gen -I ./protos ./protos/peek.proto
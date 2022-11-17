build:
# The --go_out generates the go code for protobuf
# The --go-grpc_out generates the go code for gRPC
	 protoc -I./proto --go_out=./proto/rocket/v1 \
 					  --go-grpc_out=./proto/rocket/v1 \
 					   proto/rocket/**/*.proto

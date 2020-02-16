protoc \
  --proto_path=${GOPATH}/src/github.com/gogo:. \
  --gogofaster_out=plugins=grpc:. \
  --grpc-gateway_out=grpc_api_configuration=router.yaml \  << here
  service.proto

protoc \
  --proto_path=${GOPATH}/src/github.com/gogo:. \
  --gogofaster_out=plugins=grpc:. \
  service.proto


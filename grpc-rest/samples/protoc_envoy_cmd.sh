protoc \
  -I . \
  -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.15.2/third_party/googleapis \
  --go_out=plugins=grpc:. \
  --include_imports --include_source_info \ // HL
  --descriptor_set_out=example.pb \ // HL
  ./example.proto


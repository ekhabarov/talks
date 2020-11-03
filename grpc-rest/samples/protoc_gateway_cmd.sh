protoc \
  --go_out=plugins=grpc:. \
  --grpc-gateway_out=:. \ // HL
  ./example.proto


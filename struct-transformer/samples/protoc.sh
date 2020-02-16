protoc \
  --go_out=Moptions/annotations.proto=path/to/protoc-gen-struct-transformer/options,plugins=grpc:. \
  --struct-transformer_out=package=transform:. \ // HL
  ./service.proto


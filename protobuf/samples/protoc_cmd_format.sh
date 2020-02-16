protoc \
  --proto_path=<import_dirs>:. \
  --<plugin_name>_out=<plugin_options>:<output> \
  --<plugin_name_2>_out=<plugin_options_2>:<output_2> \
  /path/to/service.proto

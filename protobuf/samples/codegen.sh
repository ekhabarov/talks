docker run --rm \
  -v /path/to/config:/config \
  -v /path/to/input:/input \
  -v /path/to/output:/output \
  apaleo/swagger-codegen generate \
  --config /config/php-config.json \
  -i /input/service.swagger.json \
  -l php \
  -o /output

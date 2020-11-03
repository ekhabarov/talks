grpcurl -plaintext \
  -d '{"name": "Boston"}' \
  127.0.0.1:5001 greeter.Hello.SayHello

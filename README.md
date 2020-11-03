# Talks

### "Protobuf-First APIs: gRPC & Co." | [Boston Golang meetup](https://www.meetup.com/bostongo/events/273933921/) | Oct 27, 2020

As a very lazy developer I try to make life as easy as possible. In an ideal
world code should write itself, but it won’t … until you say it should do this.
In this talk we will discuss code generation and how while writing gRPC API to
get REST for free.

* [slides](https://talks.godoc.org/github.com/ekhabarov/talks/grpc-rest/grpc.slide#1)

### "Interface in Go" | [Golang Winnipeg Meetup #4](https://www.meetup.com/golangwpg/events/266794339/) | Dec 4, 2019
It's a very very basic presentation about how to use interfaces, what is it and
why you need it.

* [slides](https://talks.godoc.org/github.com/ekhabarov/talks/interfaces/interfaces.slide#1)

### "protoc-gen-struct-transformer plugin" | [Gopherpalooza, 2019](http://gopherpalooza.com/) | Mountain View, CA | Nov 8, 2019

When you work with gRPC you work with an auto-generated code which contains Go
structures. If your app has a strong separation between business logic and
transport level most likely your business logins has it’s own set of structures.
The issue here is for Go two structure types with different names and even with
equal set of fields are different types, i.e. in order to publish result from BL
level to gRPC endpoint you have to convert one structure type into another,
field-by-field or you can generate such kind of transformations and I’m going to
explain how to achieve it.

* [slides](https://talks.godoc.org/github.com/ekhabarov/talks/struct-transformer/lightning.slide#1)
* [video](https://youtu.be/ifJ7enmOG9I)

### "Protobuf-first APIs: gRPC & Co." | [Golang Winnipeg Meetup #3](https://www.meetup.com/golangwpg/events/263867072/) | Sep 18, 2019

* [slides](https://talks.godoc.org/github.com/ekhabarov/talks/protobuf/protobuf.slide#1)

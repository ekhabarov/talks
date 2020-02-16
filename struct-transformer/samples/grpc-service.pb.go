import (
	"context"
)

type Order struct {
	Id     int64  `protobuf:...`
	Number string `protobuf:...`
	Status string `protobuf:...`
}

type Request struct {
	Id int64 `protobuf:...`
}

type Response struct {
	Order *Order `protobuf:...`
}

type OrderServer interface {
	Get(context.Context, Request) (*Response, error)
}


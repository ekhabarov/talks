package grpc_implementation

// OMIT
import "context" // OMIT
// OMIT
type server struct {
	svc OrderService
}

// implementation of OrderServer interface from service.pb.go

func (s *server) Get(ctx context.Context, req pb.Request) (*pb.Response, error) {
	o, err := s.svc.GetOrder(ctx, int(req.Id)) // "o" will have a type of business_logic.MyOrder // HL
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Order: &pb.Order{ // "Order" here is of type *pb.Order // HL
			Id:     int64(o.ID), // HL
			Number: o.Number,    // HL
			Status: o.Status,    // HL
		}, // HL
	}
}

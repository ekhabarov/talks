import (
	"context"
)

type (
	MyOrder struct {
		ID     int
		Number string
		Status string
	}

	OrderService interface{ GetOrder(int) (MyOrder, error) }

	server struct{ svc OrderService }
)

func (s *server) Get(ctx context.Context, req pb.Req) (*pb.Resp, error) {
	o, err := s.svc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Resp{
		Order: &pb.Order{Id: o.ID, Number: o.Number, Status: o.Status}, // HL
	}
}

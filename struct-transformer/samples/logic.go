package business_logic

// OMIT
import "context" // OMIT
// OMIT
type MyOrder struct {
	ID     int // Order message from service.pb.go has int64 type // HL
	Number string
	Status string
}

type OrderService interface {
	GetOrder(ctx context.Context, id int) (MyOrder, error)
}

// service.proto						   |	// service.pb.go // HL
option go_package = "pb";				  |	package pb
										   |
message Order {							|	type Order struct {
  int64 id = 1;							|	  Id     int64  `protobuf:...`
  string number = 2;					   |	  Number string `protobuf:...`
  string status = 3;					   |	  Status string `protobuf:...`
}										  |	}
										   |
message Request {						  |	type Request struct {
  int64 id = 1;							|	  Id int64 `protobuf:...`
}										  |	}
										   |
message Response {						 |	type Response struct {
  Order order = 1;						 |	  Order *Order `protobuf:...`
}										  |	}
										   |
service Order{							 |	type OrderServer interface {
  rpc Get (Request) returns (Response);	|	  Get(context.Context, Request) (*Response, error)
}										  |	}

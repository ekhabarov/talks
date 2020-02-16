package transform

func MyOrderToOrder(mo business_logic.MyOrder) pb.Order {
	return pb.Order{
		Id:     int64(mo.ID),
		Number: mo.Number,
		Status: mo.Status,
	}
}

func OrderToMyOrder(o pb.Order) business_logic.MyOrder {
	return business_logic.MyOrder{
		ID:     int(o.Id),
		Number: o.Number,
		Status: o.Status,
	}
}

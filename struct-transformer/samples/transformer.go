package transform

func ProductToPb(m service.Product) pb.Product {
	return pb.Product{
		Id:   m.ID,
		Name: m.Name,
	}
}

func PbToProduct(p pb.Product) service.Product {
	return service.ProductModel{
		ID:   p.Id,
		Name: p.Name,
	}
}

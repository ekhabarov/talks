package transform

// pb.Order => business_logic.MyOrder // HL
func PbToMyOrderPtr(src *pb.Order, opts ...TransformParam) *business_logic.MyOrder {...}
func PbToMyOrderPtrList(src []*pb.Order, opts ...TransformParam) []*business_logic.MyOrder {...}
func PbToMyOrderPtrVal(src *pb.Order, opts ...TransformParam) business_logic.MyOrder {...}
func PbToMyOrderPtrValList(src []*pb.Order, opts ...TransformParam) []business_logic.MyOrder {...}
func PbToMyOrderList(src []*pb.Order, opts ...TransformParam) []business_logic.MyOrder {...}
func PbToMyOrder(src pb.Order, opts ...TransformParam) business_logic.MyOrder {...}
func PbToMyOrderValPtr(src pb.Order, opts ...TransformParam) *business_logic.MyOrder {...}
func PbToMyOrderValList(src []pb.Order, opts ...TransformParam) []business_logic.MyOrder {...}

// business_logic.MyOrder => pb.Order // HL
func MyOrderToPbPtr(src *business_logic.MyOrder, opts ...TransformParam) *pb.Order {...}
func MyOrderToPbPtrList(src []*business_logic.MyOrder, opts ...TransformParam) []*pb.Order {...}
func MyOrderToPbPtrVal(src *business_logic.MyOrder, opts ...TransformParam) pb.Order {...}
func MyOrderToPbValPtrList(src []business_logic.MyOrder, opts ...TransformParam) []*pb.Order {...}
func MyOrderToPbList(src []business_logic.MyOrder, opts ...TransformParam) []*pb.Order {...}
func MyOrderToPb(src business_logic.MyOrder, opts ...TransformParam) pb.Order {...}
func MyOrderToPbValPtr(src business_logic.MyOrder, opts ...TransformParam) *pb.Order {...}
func MyOrderToPbValList(src []business_logic.MyOrder, opts ...TransformParam) []pb.Order {...}

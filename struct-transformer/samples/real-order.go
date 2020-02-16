
import "time"

type MyRealOrder struct {
	ID                int        `json:"id"`
	ShopId            string     `json:"shop_number"`
	Subtotal          float64    `json:"subtotal"`
	SubtotalTax       float64    `json:"subtotal_tax"`
	ShippingSubtotal  float64    `json:"shipping_subtotal"`
	ShippingTax       float64    `json:"shipping_tax"`
	Total             float64    `json:"total"`
	TotalTax          float64    `json:"total_tax"`
	PaymentMethod     string     `json:"payment_method"`
	ShippingMethod    string     `json:"shipping_method"`
	OrderNumber       int        `json:"order_number"`
	Discount          float64    `json:"discount"`
	Notes             string     `json:"notes"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
	PlacedAt          time.Time  `json:"placed_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	Items             []Item     `json:"order_items"`        // HL
	BillingAddress    Address    `json:"billing_address"`    // HL
	ShippingAddresses []Address  `json:"shipping_addresses"` // HL
}


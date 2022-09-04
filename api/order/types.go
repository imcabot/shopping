package order

type CompleteOrderRequest struct {
}

type CancelOrderRequest struct {
	OrderID uint `json:"order_ID"`
}

package invoice

type CreateInvoiceInput struct {
	OrderID       int     `json:"order_id"`
	Total         float64 `json:"total"`
	CustomerEmail string  `json:"customer_email"`
}

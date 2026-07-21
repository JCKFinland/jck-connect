package purchase

// PurchaseResponse represents a successful purchase response.
type PurchaseResponse struct {
	OrderReference       string `json:"order_reference"`
	TransactionReference string `json:"transaction_reference"`
	Message              string `json:"message"`
}

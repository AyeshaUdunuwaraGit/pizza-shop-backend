package models

type InvoiceItemRequest struct {
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

type CreateInvoiceRequest struct {
	CustomerName string              `json:"customer_name"`
	Items        []InvoiceItemRequest `json:"items"`
}

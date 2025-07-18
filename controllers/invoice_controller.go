package controllers

import (
	"encoding/json"
	"net/http"
	"pizza-shop-backend/config"
	"pizza-shop-backend/models"
	"time"
)

// CreateInvoice handles the creation of a complete invoice including invoice_items.
// This follows standard POS logic:
// - Each item's total = unit_price × quantity
// - Invoice total = sum of all item totals
// - Tax is applied as a percentage (e.g., 10%)
// - Net total = total + tax
func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var req models.CreateInvoiceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	tx, err := config.DB.Begin()
	if err != nil {
		http.Error(w, "Transaction start failed", http.StatusInternalServerError)
		return
	}

	var total float64 = 0
	var taxRate float64 = 0.1 // 10% VAT
	var netAmount float64 = 0

	// Step 1: Create invoice with 0 values temporarily
	var invoiceID int
	err = tx.QueryRow(
		`INSERT INTO invoices (customer_name, created_at, total_amount, tax_amount, net_amount)
		 VALUES ($1, $2, 0, 0, 0) RETURNING id`,
		req.CustomerName, time.Now(),
	).Scan(&invoiceID)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to create invoice", http.StatusInternalServerError)
		return
	}

	// Step 2: Process each item
	for _, item := range req.Items {
		var unitPrice float64
		err := tx.QueryRow("SELECT price FROM items WHERE id = $1", item.ItemID).Scan(&unitPrice)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Item not found", http.StatusBadRequest)
			return
		}

		// POS Standard: line total = unit price × quantity
		lineTotal := unitPrice * float64(item.Quantity)
		total += lineTotal

		_, err = tx.Exec(
			`INSERT INTO invoice_items (invoice_id, item_id, quantity, unit_price, total_price)
			 VALUES ($1, $2, $3, $4, $5)`,
			invoiceID, item.ItemID, item.Quantity, unitPrice, lineTotal,
		)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to insert invoice item", http.StatusInternalServerError)
			return
		}
	}

	// Step 3: Calculate tax and net amount
	taxAmount := total * taxRate
	netAmount = total + taxAmount

	// Step 4: Update invoice with totals
	_, err = tx.Exec(
		`UPDATE invoices SET total_amount = $1, tax_amount = $2, net_amount = $3 WHERE id = $4`,
		total, taxAmount, netAmount, invoiceID,
	)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to update invoice totals", http.StatusInternalServerError)
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Invoice created successfully"))
}

// GetInvoices returns a list of invoices with their items and breakdowns.
func GetInvoices(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT
			i.id, i.customer_name, i.created_at, i.total_amount, i.tax_amount, i.net_amount,
			it.name, it.category, ii.quantity, ii.unit_price, ii.total_price
		FROM invoices i
		JOIN invoice_items ii ON i.id = ii.invoice_id
		JOIN items it ON it.id = ii.item_id
		ORDER BY i.id DESC
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch invoices", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Item struct {
		ItemName   string  `json:"item_name"`
		Category   string  `json:"category"`
		Quantity   int     `json:"quantity"`
		UnitPrice  float64 `json:"unit_price"`
		TotalPrice float64 `json:"total_price"`
	}

	type Invoice struct {
		InvoiceID    int       `json:"invoice_id"`
		CustomerName string    `json:"customer_name"`
		CreatedAt    time.Time `json:"created_at"`
		TotalAmount  float64   `json:"total_amount"`
		TaxAmount    float64   `json:"tax_amount"`
		NetAmount    float64   `json:"net_amount"`
		Items        []Item    `json:"items"`
	}

	invoiceMap := make(map[int]*Invoice)

	for rows.Next() {
		var id int
		var customerName string
		var createdAt time.Time
		var totalAmount, taxAmount, netAmount float64
		var item Item

		err := rows.Scan(&id, &customerName, &createdAt, &totalAmount, &taxAmount, &netAmount,
			&item.ItemName, &item.Category, &item.Quantity, &item.UnitPrice, &item.TotalPrice)
		if err != nil {
			http.Error(w, "Failed to scan invoice data", http.StatusInternalServerError)
			return
		}

		if _, exists := invoiceMap[id]; !exists {
			invoiceMap[id] = &Invoice{
				InvoiceID:    id,
				CustomerName: customerName,
				CreatedAt:    createdAt,
				TotalAmount:  totalAmount,
				TaxAmount:    taxAmount,
				NetAmount:    netAmount,
				Items:        []Item{},
			}
		}
		invoiceMap[id].Items = append(invoiceMap[id].Items, item)
	}

	var result []Invoice
	for _, invoice := range invoiceMap {
		result = append(result, *invoice)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

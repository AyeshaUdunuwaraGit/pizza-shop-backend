package routes

import (
	"net/http"
	"pizza-shop-backend/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.GetAllItems(w, r)
		} else if r.Method == http.MethodPost {
			controllers.CreateItem(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/invoices", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.CreateInvoice(w, r)
		} else if r.Method == http.MethodGet {
			controllers.GetInvoices(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

# 🍕 Pizza Shop Billing System - Go Backend

This project is a simple **Pizza Shop Billing System** backend built using **Go (Golang)**, **PostgreSQL**, and follows a clean **MVC structure**. It allows managing items (pizzas, toppings, beverages), creating invoices, and viewing all invoices with item breakdown.

---

## 🛠 Tech Stack

- **Backend**: Go (net/http)
- **Database**: PostgreSQL
- **Architecture**: MVC (Model-View-Controller)
- **Tools**: Postman (for testing APIs)

---

## 📁 Project Structure## 

pizza-shop-backend/
├── config/           # DB connection
│ └── db.go
├── controllers/      # Business logic
│ ├── invoice_controller.go
│ └── item_controller.go
├── models/           # Request/response models
│ ├── invoice.go
│ └── item.go
├── routes/           # Route definitions
│ └── routes.go
├── main.go           # Entry point
├── go.mod
├── go.sum
└── README.md         # This file

yaml
Copy
Edit

---

# 🔌 API Features

## 1. **Item Management**

- **GET** `/items` → List all available items (pizzas, toppings, beverages)
- **POST** `/items` → Add a new item

**Item JSON Example:**

{
  "name": "Spicy Devilled Beef Pizza",
  "price": 1400,
  "category": "Pizza"
}

## 2.  **Invoice Management**
POST /invoices → Create a new invoice

GET /invoices → List all invoices with breakdown

Invoice Creation Request:

json
Copy
Edit
{
        "invoice_id": 1,
        "customer_name": "Nimal Perera",
        "created_at": "2025-07-19T00:42:38.00956Z",
        "total_amount": 3500,
        "tax_amount": 350,
        "net_amount": 3850,
        "items": [
            {
                "item_name": "Classic Sri Lankan Chicken Curry Pizza",
                "category": "Pizza",
                "quantity": 2,
                "unit_price": 1200,
                "total_price": 2400
            },
            {
                "item_name": "Vegetarian Pol Sambol Pizza",
                "category": "Pizza",
                "quantity": 1,
                "unit_price": 1100,
                "total_price": 1100
            }
        ]
    }


### Invoices are calculated using standard POS logic: 

Line Total = unit_price × quantity

Subtotal = sum of line totals

Tax (10%) is added to get the final total

🧾 Sample Invoice Response

[
  {
    "invoice_id": 3,
    "customer_name": "Ayesha Udunuwara",
    "created_at": "2025-07-19T00:59:56Z",
    "item_name": "Spicy Devilled Beef Pizza",
    "category": "Pizza",
    "quantity": 2,
    "unit_price": 1400,
    "total_price": 2800
  }
]
Note: This response shows invoice line items. You may enhance this by grouping items under invoice headers if needed.

### 3.  **🧪 How to Run & Test**

Clone this repo

git clone https://github.com/AyeshaUdunuwaraGit/pizza-shop-backend.git
cd pizza-shop-backend
Configure database in config/db.go

connStr := "host=localhost port=5432 user=postgres password=@1234 dbname=pizza_shop sslmode=disable"

### Run the Go server 

go run main.go

## Test using Postman or curl ##

GET http://localhost:8080/items

POST http://localhost:8080/items

GET http://localhost:8080/invoices

POST http://localhost:8080/invoices

## ✅ Assessment Requirements Fulfilled 

 Add new pizzas, toppings, and beverages via API

 Create invoice with customer name, selected items, quantities

 Apply tax, calculate line totals, subtotals, net totals (POS standard)

 Retrieve invoice list with detailed item breakdown

 Organized using MVC structure

 Clear route structure and separation of concerns

 Database integration with PostgreSQL

 Code testable via Postman

## 📌 Notes
Item categories must be one of: Pizza, Topping, or Beverage

You can enhance the project by adding:

Authentication

Invoice grouping by header

Discount and coupon support

Unit tests using Go's testing framework

## 📞 Contact
Author: Ayesha Udunuwara
Email: a.n.s.udunuwara@gmail.com
Feel free to reach out for any queries or suggestions!
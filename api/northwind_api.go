package api

import (
	"fmt"
	"net/http"

	"github.com/headwinds/northwind-frostpunk/database"

	"github.com/headwinds/northwind-frostpunk/api/controllers/orders"
	"github.com/headwinds/northwind-frostpunk/api/controllers/products"
)

func NorthwindApi() {
	db := database.Connect();

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Northwind Frostpunk API v0.0.1")
    })
	
	// Products
	productsHandler := products.ProductsController(db)
	http.HandleFunc("/products", productsHandler.GetProducts)

	// Orders
	ordersHander := orders.OrdersController(db)
	http.HandleFunc("/orders", ordersHander.GetOrders)

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "8080"),
	}

	s.ListenAndServe()
}
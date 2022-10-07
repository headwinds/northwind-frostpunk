package api

import (
	"fmt"
	"net/http"

	"github.com/headwinds/northwind-frostpunk/database"

	"github.com/headwinds/northwind-frostpunk/api/controllers/customer_journey"
	"github.com/headwinds/northwind-frostpunk/api/controllers/game"
	"github.com/headwinds/northwind-frostpunk/api/controllers/orders"
	"github.com/headwinds/northwind-frostpunk/api/controllers/products"
)

func NorthwindApi() {
	db := database.Connect();

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "/")
    })

	// GAME 
	startGameHandler := game.StartGameController(db)
	http.HandleFunc("/game/start", startGameHandler.StartGame)

	// Products
	productsHandler := products.ProductsController(db)
	http.HandleFunc("/products", productsHandler.GetProducts)

	// Orders
	ordersHander := orders.OrdersController(db)
	http.HandleFunc("/orders", ordersHander.GetOrders)

	// Customer Journey
	customerJourneyHander := customer_journey.CustomerJourneyController(db)
	http.HandleFunc("/journies", customerJourneyHander.GetCustomerJournies)

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "8080"),
	}

	s.ListenAndServe()

	fmt.Println("+ Northwind Frostpunk API v0.0.1 +")
}
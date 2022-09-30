package orders

/*
	ok copilot what do we want to do to practice our API skills?
	- create a new order
	- get an order
	- update an order
	- delete an order
	- get all orders
	- get all orders for a customer
*/

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpResp struct{
    Status      int         `json:"status"`
    Description string      `json:"description"`
    Body        interface{} `json:"body"`
}

type Order struct{
    OrderId		string  `json:"order_id"`
    CustomerId	string  `json:"customer_id"`
    ShipCity	string  `json:"ship_city"`
}

type DatabaseHandler struct {
	db *sql.DB
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func OrdersController(db *sql.DB) *DatabaseHandler {
	return &DatabaseHandler{
		db: db,
	}
}

func (h *DatabaseHandler) GetOrders(w http.ResponseWriter, r *http.Request){
	if err := h.db.Ping(); err != nil {
		fmt.Println("DB Error")
	}

	var orders []Order
	rows, err := h.db.Query(`SELECT "order_id", "customer_id", "ship_city" FROM "orders" LIMIT 3`)
	checkErr(err)

	defer rows.Close()
	for rows.Next() {

		var order Order
		err = rows.Scan(&order.OrderId, &order.CustomerId, &order.ShipCity)
		checkErr(err)
	
		orders = append(orders, order)
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(HttpResp{Status: 200, Body: orders})

	checkErr(err)

}

func (h *DatabaseHandler) SimulateOrders(w http.ResponseWriter, r *http.Request){
	if err := h.db.Ping(); err != nil {
		fmt.Println("DB Error")
	}

	var orders []Order
	rows, err := h.db.Query(`SELECT "order_id", "customer_id", "ship_city" FROM "orders" LIMIT 3`)
	checkErr(err)

	defer rows.Close()
	for rows.Next() {

		var order Order
		err = rows.Scan(&order.OrderId, &order.CustomerId, &order.ShipCity)
		checkErr(err)
	
		orders = append(orders, order)
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(HttpResp{Status: 200, Body: orders})

	checkErr(err)

}
package products

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type HttpResp struct{
    Status      int         `json:"status"`
    Description string      `json:"description"`
    Body        interface{} `json:"body"`
}

type Product struct{
    ProductName		string  `json:"product_name"`
    UnitPrice		float64 `json:"unit_price"`
    UnitsInStock	int  	`json:"units_in_stock"`
}

type DatabaseHandler struct {
	db *sql.DB
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ProductsController(db *sql.DB) *DatabaseHandler {
	return &DatabaseHandler{
		db: db,
	}
}

func (h *DatabaseHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Ping(); err != nil {
		fmt.Println("DB Error")
	}

	var products []Product
	rows, err := h.db.Query(`SELECT "product_name", "unit_price", "units_in_stock" FROM "products" LIMIT 3`)
	checkErr(err)

	defer rows.Close()
	
	for rows.Next() {
		var product Product
	
		err = rows.Scan(&product.ProductName, &product.UnitPrice, &product.UnitsInStock)
		checkErr(err)
	
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(HttpResp{Status: 200, Body: products})

	checkErr(err)

}

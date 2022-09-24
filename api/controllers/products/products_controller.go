package products

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

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

	rows, err := h.db.Query(`SELECT "product_name", "unit_price", "units_in_stock" FROM "products" LIMIT 3`)
	checkErr(err)

	defer rows.Close()
	for rows.Next() {
		var product_name string
		var unit_price float64
		var units_in_stock int
	
		err = rows.Scan(&product_name, &unit_price, &units_in_stock)
		checkErr(err)
	
		fmt.Println(product_name, unit_price, units_in_stock)
	}

	checkErr(err)
}

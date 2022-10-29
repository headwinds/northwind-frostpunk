package products

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/headwinds/northwind-frostpunk/api/types"

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

	fmt.Println("GetProducts")

	pg := r.URL.Query().Get("page")

	page, err := strconv.Atoi(pg)
	if err != nil {
			page=1
	}

	iPP := r.URL.Query().Get("itemsPerPage")

	itemsPerPage, err := strconv.Atoi(iPP)
	if err != nil {
		itemsPerPage=10
	}

	limit := int(itemsPerPage)
	offset := limit * (int(page) - 1)

	var products []types.Product
	rows, err := h.db.Query(`SELECT "product_name", "unit_price", "units_in_stock" FROM "products" LIMIT ` + strconv.Itoa(limit) + ` OFFSET ` + strconv.Itoa(offset))
	checkErr(err)

	defer rows.Close()
	
	for rows.Next() {
		var product types.Product
	
		err = rows.Scan(&product.ProductName, &product.UnitPrice, &product.UnitsInStock)
		checkErr(err)
	
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(types.ProductsHttpResp{Status: 200, Description: "Hey headwinds", Body: products})

	checkErr(err)

}




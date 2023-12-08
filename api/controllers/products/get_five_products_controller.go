package products

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/headwinds/northwind-frostpunk/api/types"
	"github.com/headwinds/northwind-frostpunk/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
)

// Get all products
// response and request handlers
func GetFiveProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Create database connection
	connPool, err := pgxpool.NewWithConfig(context.Background(), database.Config())
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}

	defer connPool.Close()

	// before proceeding, ping the database if it is available
	err = connPool.Ping(ctx)
	if err != nil {
		log.Fatal("Error while pinging the database!!")
		jsonErrorResponse := types.JsonMessageResponse{Type: "Errror", Message: "No connection available"}
		json.NewEncoder(w).Encode(jsonErrorResponse)
	}

	fmt.Println("Connected to the database!!")
	query := `SELECT * FROM products LIMIT 5`

	rows, err := connPool.Query(context.Background(), query)

	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[types.Product])
	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		return
	}
	for _, p := range products {
		fmt.Printf("%s\n", p.ProductName)
	}
	jsonDataResponse := types.JsonDataResponse{Type: "Success", Message: "Products fetched successfully", Data: products}

	json.NewEncoder(w).Encode(jsonDataResponse)
}

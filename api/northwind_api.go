package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/headwinds/northwind-frostpunk/database"
	//"github.com/headwinds/northwind-frostpunk/api/controllers/customer_journey"
	//"github.com/headwinds/northwind-frostpunk/api/controllers/game"
	//"github.com/headwinds/northwind-frostpunk/api/controllers/orders"
	//"github.com/headwinds/northwind-frostpunk/api/controllers/products"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Main function
func NorthwindApi() {

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints
	router.HandleFunc("/", GetHome).Methods("GET")

	// Get all the products
	router.HandleFunc("/products", GetProducts).Methods("GET")
	//router.HandleFunc("/products", GetProducts).Methods("GET")
	//router.HandleFunc("/products", CreateProducts).Methods("POST")

	// serve the app
	fmt.Println("Northwind Frostpunk at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Product struct {
	ProductId       int     `db:"product_id"`
	ProductName     string  `db:"product_name"`
	SupplierId      int     `db:"supplier_id"`
	CategoryId      int     `db:"category_id"`
	QuantityPerUnit string  `db:"quantity_per_unit"`
	UnitPrice       float64 `db:"unit_price"`
	UnitsInStock    int     `db:"units_in_stock"`
	UnitsonOrder    int     `db:"units_on_order"`
	ReorderLevel    int     `db:"reorder_level"`
	Discontinued    int     `db:"discontinued"`
}

type JsonResponse struct {
	Type    string    `json:"type"`
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}

// Get all products
// response and request handlers
func GetProducts(w http.ResponseWriter, r *http.Request) {

	// Create database connection
	connPool, err := pgxpool.NewWithConfig(context.Background(), database.Config())
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}

  defer connPool.Close()

	fmt.Println("Connected to the database!!")
	query := `SELECT * FROM products LIMIT 5`

	rows, err := connPool.Query(context.Background(), query)

	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[Product])
	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		return
	}
	for _, p := range products {
		fmt.Printf("%s\n", p.ProductName)
	}

	json.NewEncoder(w).Encode(products)
}

// util help logging
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("ELEPHANT_CONNECTION_STR"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Elephant never forgets!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	hello := "Northwind Frostpunk at your service!"

	fmt.Fprintf(w, hello)
	fmt.Println(hello)
}

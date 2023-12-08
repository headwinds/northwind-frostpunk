package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
  "github.com/headwinds/northwind-frostpunk/api/types"
	//"github.com/headwinds/northwind-frostpunk/api/controllers/customer_journey"
	"github.com/headwinds/northwind-frostpunk/api/controllers/game"
	//"github.com/headwinds/northwind-frostpunk/api/controllers/orders"
	"github.com/headwinds/northwind-frostpunk/api/controllers/products"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Main function
func NorthwindApi() {

	// Init the mux router
	router := mux.NewRouter()

  ctx := context.Background()
  connPool, err := pgxpool.New(ctx, os.Getenv("ELEPHANT_CONNECTION_STR"))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
    os.Exit(1)
  }
  defer connPool.Close()

  // before proceeding, ping the database if it is available
  err = connPool.Ping(ctx)
  if err != nil {
    log.Fatal("Error while pinging the database!!")
  }

	// Home
	router.HandleFunc("/", GetHome).Methods("GET")

	// Products
  router.HandleFunc("/products/five", products.GetFiveProducts).Methods("GET")
  router.HandleFunc("/products", products.GetProducts).Methods("GET")

  // Game 
  router.HandleFunc("/game/start", game.StartGame).Methods("GET")
  router.HandleFunc("/game/turn/next", game.NextTurn).Methods("GET")

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

// util help logging
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	hello := "Northwind Frostpunk at your service!"
	fmt.Println(hello)
  jsonMessage := types.JsonMessageResponse{Type: "Success", Message: hello}
  json.NewEncoder(w).Encode(jsonMessage)
}

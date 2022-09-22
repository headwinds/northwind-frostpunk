package main

import (
	"database/sql"
	"fmt"

	"github.com/headwinds/northwind-frostpunk/hello"

	_ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 55432
  user     = "postgres"
  password = "postgres"
  dbname   = "northwind"
)

func main() {

  message := hello.Hello()
  fmt.Println(message)

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")

  getProducts(db, 0, 10)
}

// thank you Github copilot for this snippet
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getProducts(db *sql.DB, start, count int) {
	rows, err := db.Query(`SELECT "product_name", "unit_price", "units_in_stock" FROM "products"`)
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
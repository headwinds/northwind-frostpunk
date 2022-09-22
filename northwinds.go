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
}
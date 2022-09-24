package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 55432
  user     = "postgres"
  password = "postgres"
  dbname   = "northwind"
)

func Connect() *sql.DB {
  db := connectToPostgres();
  return db
}

func connectToPostgres() *sql.DB {

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")

  return db
}

package database

import (
	"database/sql"
	"fmt"
	//"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	//"github.com/spf13/viper"
)

const (
	host     = "localhost"
	port     = 55432
	user     = "postgres"
	password = "postgres"
	dbname   = "northwind"
)

func Connect() *sql.DB {
	db := connectToPostgres()
	return db
}

// use viper package to read .env file
// return the value of the key
// https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
/*
On repl, use the secrets approach instead
func viperEnvVariable(key string) string {

  // SetConfigFile explicitly defines the path, name and extension of the config file.
  // Viper will use this and not check any of the config paths.
  // .env - It will search for the .env file in the current directory
  viper.SetConfigFile(".env")

  // Find and read the config file
  err := viper.ReadInConfig()

  if err != nil {
    log.Fatalf("Error while reading config file %s", err)
  }

  // viper.Get() returns an empty interface{}
  // to get the underlying type of the key,
  // we have to do the type assertion, we know the underlying value is string
  // if we type assert to other type it will throw an error
  value, ok := viper.Get(key).(string)

  // If the type is a string then ok will be true
  // ok will make sure the program not break
  if !ok {
    log.Fatalf("Invalid type assertion")
  }

  return value
}
*/

func connectToPostgres() *sql.DB {

	//elephantSqlConnectionStr := viperEnvVariable("DBURL")
	//str := viperEnvVariable("IS_ELEPHANTSQL")
	elephantSqlConnectionStr := os.Getenv("DBURL")
	str := os.Getenv("IS_ELEPHANTSQL")

	fmt.Println("Secrets: ", elephantSqlConnectionStr)

	isElephantSql, _ := strconv.ParseBool(str)
	//isElephantSql := viperEnvVariable("IS_ELEPHANTSQL")

	//fmt.Println("viperenv", elephantSqlConnectionStr)

	// Find and read the config file
	/*
	  err := viper.ReadInConfig()

	  if err != nil {
	    log.Fatalf("Error while reading config file %s", err)
	  }
	*/

	// connect with struct or connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var connStr string

	if isElephantSql {
		connStr = elephantSqlConnectionStr
	} else {
		fmt.Println("not elephant")
		connStr = psqlInfo
	}

	connStr = elephantSqlConnectionStr //"postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	if isElephantSql {
		fmt.Println("Successfully connected to the remote Northwind database on ElephantSQL!")
	} else {
		fmt.Println("Successfully connected to the local Northwind database")
	}

	return db
}

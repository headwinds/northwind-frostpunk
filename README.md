# northwind-frostpunk
<img src="frostpunk.png" />

This is a study exploring golang, the northwind database, and an order status story.  

1. connect Go to Postgres featuring the Northwind database augmented with Frostpunk tables
2. interact with a Go CLI 
3. simulate orders, view a report & have fun!

Golang 1.8+ [install gvm](https://jimkang.medium.com/install-go-on-mac-with-homebrew-5fa421fc55f5)

## Example SQL Query 

``` 
 SELECT  C.category_name
       ,COUNT(P.category_id) AS total_products
  FROM categories AS C
  JOIN products AS P
    ON P.category_id = C.category_id
 GROUP BY C.category_name
 ORDER BY total_products DESC;
```

## 1. Database

First, start the up the database container in a terminal using the [northwind_psql](https://github.com/pthom/northwind_psql) repo not this repo! 

```
cd northwind_psql
docker-compose up
```

From the pgAdmin4 or psql, you will need to modify this database to include the frostpunk-inspired tables.

## 2. API

Open a second terminal for the API which is back to this repo. 

```
go run . 
```

You should see the message `Successfully connected to the Northwind database` 

## 3. CLI

The API has to be running in a separate terminal before you can interact with it via the CLI

Open a third terminal to run the CLI

```
cd cli
go run . 
```
## Tests

```
go test ./...
```

## Log

[Captain's Log](log/captainsLog.md)


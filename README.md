# northwind-frostpunk
<img src="frostpunk.png" />

This is a study exploring golang, the northwind database, and an order status story.  

1. connect Go to Postgres featuring the Northwind database
2. interact with a Go CLI 
3. view a report 

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
## Log

[Captains Log](log/captainsLog.md)


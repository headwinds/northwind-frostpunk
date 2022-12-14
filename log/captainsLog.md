# northwind-frostpunk
a tutorial exploring golang, the northwind database, and the post apocalypse 

This tutorial is geared towards Mac users.

why am I making this? I've become familar with the Northwind database to teach SQL but I thought many of the tutorials I went through to be crudely designed and I simply desired something that would encourage me to explore and expand upon it. 

I also wanted to work on Golang again since the last app, Tank. I built Tank (think aquarium) to produce a daily report (crontab) looking at our feeds folder and would look into a dozen third-party feed folders. It woud count the incoming, proceeded and error files and attempt to raise an insight or call to action. Since it's been several months since I last touched it, it feels like I've forgotten all the syntax. Fortunately, building software is like riding bike! Even if it is a bit rickety at the start...

<img src="../frostpunk.png" />

## Day 1 Overview

1. setup and connect to postgres
2. learn a bit of golang 
3. connect golang to postgres
4. explore the database and learn about CTEs
5. play with the data & time inspired by [Frostpunk](https://en.wikipedia.org/wiki/Frostpunk)!

## PostgreSQL 

[northwind_psql](https://github.com/pthom/northwind_psql) is a github repo that sets up Postgres with a docker container. 

In addition to its readme, I would add that the following guide post that once the docker is up and finishes installing the database, open another another terminal window and test the conection.  

```
docker-compose exec db psql -U postgres -d northwind
```

After you've tested the connection, you can also try to connect with a tool like [pgAdmin4](https://www.pgadmin.org/) or [DBeaver](https://dbeaver.io/); they are both excellent! 

#### Connection 

```
host: localhost
username: postgres 
password: postgres 
dname: northwind 
port: 55432 
```

Initialy, I found it challenging to connect to the database running in its docker and this is connection details that I ended up using by trial and error.

## Golang

I followed this golang [getting-started](https://go.dev/doc/tutorial/getting-started) tutorial as a refresher since I hadn't touched golang in several months.

Next, after I was to output some pithy text and learned the `go mod tidy` command, I tackled this [go and postgres](https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/) tutorial. 

As I already wrestled the postgres setup, the code provided in the tutorial worked right away. Now, I wanted to import my own function into my main project so I could log a pithy quote. I reviewed the [go mod](https://go.dev/doc/tutorial/create-module) tutorial and converted my original hello.go script into a function that I could import.  

#### Log ????

-- Sept 21/2022 -- 

Day 1 took me about 4 hours in 2 sessions

As I accomplished the first 3 steps in my plan, I think I'll needed a couple more hours to put the finishing touches on this guided tutorial. I still need to:

4. explore the database and learn about CTEs
5. play with the data & time inspired by [Frostpunk](https://en.wikipedia.org/wiki/

Anyways, I'm happy with this progress [as the hours pass...](https://www.youtube.com/watch?v=RQBDSciMe8c)

## Day 2 Survive the night! 

<img src ='https://raw.githubusercontent.com/pthom/northwind_psql/master/ER.png' />

A database mapped out in [UML](https://www.microsoft.com/en-ca/microsoft-365/business-insights-ideas/resources/guide-to-uml-diagramming-and-database-modeling) is an adventure for your eyes. 

Let's begin day 2 will pulling a few examples from this [Products Table Exercises](https://www.w3resource.com/mysql-exercises/northwind/products-table-exercises/) tutorial. 

I like how this author has included a [relational algebra](https://www.javatpoint.com/dbms-relational-algebra) tree diagram. 

<img src="https://www.w3resource.com/w3r_images/mysql-products-table-exercises-relational-algebra-tree-diagram-10.png" />

I'm not the alone in this desire to [learn SQL](https://sqlbolt.com/) and [master its style](https://www.sqlstyle.guide/). Musili Adebayo has prepared an excellent [50 query guide](https://musiliadebayo.medium.com/50-sql-practice-queries-and-answers-3fc896650b2e).  

Over on stackoverflow, this user asks for [more](https://stackoverflow.com/questions/1065652/looking-for-exercises-to-learn-sql-using-the-northwind-database). 

https://itnext.io/go-tutorial-database-queries-on-go-with-postgresql-part-iii-c8907729c2f

4. Write a query to get most expense and least expensive Product list (name and unit price).

```
SELECT product_name, unit_price 
FROM products 
ORDER BY unit_price DESC;
```

6. Write a query to get Product list (id, name, unit price) where products cost between $15 and $25.

```
SELECT product_name, unit_price
FROM products
WHERE (((unit_price)>=15 And (unit_price)<=25) 
AND ((products.discontinued)=1))
ORDER BY products.unit_price DESC;
```

9. Write a query to count current and discontinued products.

One of favourite queries is the ability to simply count and roll up records. The Count function is the beginning of your data aggregation journey! 

```
SELECT Count(product_name)
FROM products
GROUP BY discontinued;
```

When dealing with physical products, there is a real sense of scarcity since products may go out of stock. We can't sell what we don't have.

10. Write a query to get Product list (name, units on order , units in stock) of stock is less than the quantity on order.
```
SELECT product_name,  units_on_order, units_in_stock
FROM products
WHERE (((discontinued)=0) AND ((units_in_stock)<units_on_order));
```
Next, I looked at this tutorial [Northwind Queries Part 1](https://www.geeksengine.com/database/problem-solving/northwind-queries-part-1.php) from Geeks Engine 


4. This is a rather simple query to get an alphabetical list of products.

```
SELECT DISTINCT b.*, a.category_name
FROM categories a 
INNER JOIN products b on a.category_id = b.category_id
WHERE b.discontinued = 0
ORDER BY b.product_name;
```

11. Products Above Average Price

This query shows how to use sub-query to get a single value (average unit price) that can be used in the outer-query.

```
SELECT DISTINCT product_name, unit_price
FROM products
WHERE unit_price > (SELECT AVG(unit_price) FROM products)
ORDER BY unit_price;
```

14. Quarterly Orders by Product

This query shows how to convert order dates to the corresponding quarters. It also demonstrates how SUM function is used together with CASE statement to get sales for each quarter, where quarters are converted from OrderDate column.

```
SELECT a.product_name, 
    d.company_name, 
    EXTRACT(YEAR FROM CAST(c.order_date AS DATE)) as order_year,
    (sum(case EXTRACT (QUARTER FROM c.order_date) when '1' 
        then b.unit_price*b.quantity*(1-b.discount) else 0 end), 0) "Qtr 1",
    (sum(case EXTRACT (QUARTER FROM c.order_date) when '2' 
        then b.unit_price*b.quantity*(1-b.discount) else 0 end), 0) "Qtr 2",
    (sum(case EXTRACT (QUARTER FROM c.order_date) when '3' 
        then b.unit_price*b.quantity*(1-b.discount) else 0 end), 0) "Qtr 3",
    (sum(case EXTRACT (QUARTER FROM c.order_date) when '4' 
        then b.unit_price*b.quantity*(1-b.discount) else 0 end), 0) "Qtr 4" 
from products a 
inner join order_details b on a.product_id = b.product_id
inner join orders c on c.order_id = b.order_id
inner join customers d on d.customer_id = c.customer_id 
where c.order_date between date('1997-01-01') and date('1997-12-31')
group by a.product_name, 
    d.company_name, 
    EXTRACT(YEAR FROM CAST(c.order_date AS DATE))
order by a.product_name, d.company_name;
```
There is stretch goal that I need to look since the formatting of the numbers looks funny. 

I know that I need to use either `char_at` or the `FORMAT` function to improve it.  

Searching for more sources of northwind practice, I stumbled across [this github repo](https://github.com/emnasc/northwind-sql-practice) by 
Emanoel Nascimento where the challenge became can I create the questions if I had solutions! 

I grabbed a few of his examples from the [sql](https://github.com/emnasc/northwind-sql-practice/blob/master/sql-practice-execises.sql) he provided. 

20.
``` 
 SELECT  C.category_name
       ,COUNT(P.category_id) AS total_products
  FROM categories AS C
  JOIN products AS P
    ON P.category_id = C.category_id
 GROUP BY C.category_name
 ORDER BY total_products DESC;
```
29.
```
SELECT  ORD.employee_id
       ,EMP.last_name
	   ,ODE.order_id
	   ,PRO.product_name
	   ,ODE.quantity
  FROM order_details AS ODE
  JOIN orders AS ORD
    ON ORD.order_id = ODE.order_id
  JOIN employees AS EMP
    ON ORD.employee_id = EMP.employee_id
  JOIN products AS PRO
    ON PRO.product_id = ODE.product_id
 ORDER BY ODE.order_id ASC; 
```
30.

```
SELECT  C.customer_id AS customer_customer_id
	   ,O.customer_id AS orders_customer_id
  FROM customers AS C
  LEFT JOIN orders AS O
	ON C.customer_id = O.customer_id
 WHERE O.customer_id IS NULL;
```
What happens if we want the customer to not be null?
```
SELECT  C.customer_id AS customer_customer_id
	   ,O.customer_id AS orders_customer_id
  FROM customers AS C
  LEFT JOIN orders AS O
	ON C.customer_id = O.customer_id
 WHERE O.customer_id IS NOT NULL;
```

While it is possible to [add the images](http://www.geeksengine.com/article/export-access-to-mysql-5.html) or even have [AI generate them](https://generated.photos/), since this geared towards a text-based adventure they are not necessary and rather poor photography anyways. Perhaps, an AI 

## Frostpunk

Besides [3 minute chess](https://lichess.org/) or [Wordle](https://www.nytimes.com/games/wordle/index.html), I haven't played a serious computer game since I gave up Starcraft 2 a few years ago. I wanted forever it seems to finally play Frostpunk which is [now 75% off](https://store.steampowered.com/agecheck/app/323190/) on steam.

#### Hope

```
CREATE TABLE employee_hope
(
  employee_id integer NOT NULL,
  hope integer NOT NULL,
  hunger integer NOT NULL,
  drive integer NOT NULL,
  grant_date timestamp default current_timestamp,
  CONSTRAINT employee_hope_id_fkey FOREIGN KEY (employee_id)
      REFERENCES employee (employee_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);
```

#### Profession 

```
CREATE TABLE profession(
  profession_id serial primary key,
  profession_name varchar(255) UNIQUE NOT NULL,
  profession_salary integer NOT NULL
);
```
```
INSERT INTO profession (profession_name, profession_salary)
VALUES
    ('Developer', 120000),
    ('Trader', 80000),
    ('Manager', 140000);
```

```
CREATE TABLE employee_profession
(
  employee_id integer NOT NULL,
  profession_id integer NOT NULL,
  grant_date timestamp default current_timestamp,
  PRIMARY KEY (employee_id, profession_id),
  CONSTRAINT employee_profession_profession_id_fkey FOREIGN KEY (profession_id)
      REFERENCES profession (profession_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT employee_profession_employee_id_fkey FOREIGN KEY (employee_id)
      REFERENCES employee (employee_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);
```

In her [learning go by exmample](https://dev.to/aurelievache/learning-go-by-examples-part-2-create-an-http-rest-api-server-in-go-1cdm) series, Aurelie Vache has created an impressive set. Her gameboy one is especially fun but for this guide I'll leverage her post on building the RESTful endpoints and the CLI! 

It's always nice to reflective on work from many moons and wonder what could have been. Looking back on this [chatbot customer journey](https://observablehq.com/@headwinds/decision-tree) or [hiring decision](https://observablehq.com/d/03897904957ef390) is standing-on-our-shoulders-of-giants development where you need to tweak the sankey diagram to it meet your needs. 

This reflection was possibly more procrastination as it took me away from my CLI goal but wanted to acknowledge as flight of interest and remind that it's out there and could assist with current issues at work where we are struggling to track order status.  

After all the database work, I ran out of juice by the end of the session so ended it with one last push (with the help of this [Golang Postgres](https://hevodata.com/learn/golang-postgres/) post) to actually hook up the query to go and simply print a list of products to the terminal.

<img src="terminal_day2.png" />

#### Log ???? 

-- Sept 22/2022 -- 

I had less time today than yesterday with work heating up. At a large enterprise, at times we only have 60% of a solution ready to serve our trading team and call centre. Developers often need to refer to SOP (service operationg procedures) that document how to run scripts in our admin tool to patch problems where a bit of UI and a button could have saved the day to complete the feature. 

As someone who cares about product and fullstack development, there is where I see large gaps in our UI and the services we offer.  Obviously, I'm proud of my own SOP contributions. 

This write up was banged out to [Nina Las Vegas b2b Swick DJ Set | Keep Hush Live](https://www.youtube.com/watch?v=8d0mHipC1LA) and [HAAi B2B Daniel Avery | Boiler Room x Glitch Festival 2022](https://www.youtube.com/watch?v=WFexmmlpXe0).

## Day 3 

- Points, gamification & order status
- [poslog](https://www.yellowfinbi.com/blog/2013/06/yfcommunitynews-poslog-one-data-container-to-rule-them-all-well-at-least-for-retail-141038/)
- [Composition](https://medium.com/bitaksi-tech/object-oriented-approach-in-go-179c85486c76) 

#### Order Status

```
CREATE TABLE order_status(
  order_status_id serial primary key,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);
```
```
ALTER TABLE order_status 
    ADD COLUMN order_id INTEGER 
    REFERENCES orders (order_id);
```  

OR in one go....

```
DROP TABLE order_status;
```

```
CREATE TABLE order_status(
  order_status_id serial primary key,
  order_status varchar(255) NOT NULL,
  order_id INT,
   CONSTRAINT fk_order
      FOREIGN KEY(order_id) 
	  REFERENCES orders(order_id),
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);
```


What happens to the database over 10 days?

What are 3 problems that the developer could face?

- order is in the run status via human coding error
- fraud
- business wants a report on revenue to see top amd bottom sellig products

I'm picking up more [syntax patterns in Golang](https://medium.com/@kdnotes/golang-naming-rules-and-conventions-8efeecd23b68) as well as expanding my import strategy to have more a deeply nested folder structure. 

In order to import a package, we reference the functions by using the package name. 

For instance, I have a folder: /api/orders/orders_service.go 

After I import this folder into a parent file:

```
import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/headwinds/northwind-frostpunk/api/orders"
)
```

I could then use any function within the orders folder like `orders.OrderService()`. 

So now that I have controllers managing my routes I need a way to [access the database connection](https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/). 

After refactoring my controllers based on that post, I ran into an issue where my database connection was closed when I tried to use it my routes. I learned that I needed to remove `defer db.Close()` from this nice [stackoverflow](https://stackoverflow.com/questions/57867635/mysql-database-gets-error-in-golang-sql-database-is-closed) and that Golang does an excellent job of managing closingn the connection. 

Now that I could query the database and print rows in the terminal I want to return the results as json like any proper RESTful so down [another rabbit hole](https://stackoverflow.com/questions/45914863/how-to-return-json-data-without-escape-string-in-restful-api-on-golang) I go to learn more about http responses and then I was able to my `/orders` endpoint to return: 

```
{"status":200,"description":"","body":[{"order_id":"10248","customer_id":"VINET","ship_city":"Reims"},{"order_id":"10249","customer_id":"TOMSP","ship_city":"M??nster"},{"order_id":"10250","customer_id":"HANAR","ship_city":"Rio de Janeiro"}]}
```

<img src="orders_response.png" />

Design patterns are extremely valuable in programming. When you approach another language, you can lean on these patterns to help you navigate and build workflows. If a language does not support common patterns, it is sign that the language may not be robust and mature as other langauges. 

One of my favourite patterns is [Functional Programming](https://fsharpforfunandprofit.com/fppatterns/), and when coing in Java, Javascript, Go or Python, I attempt to use these patterns as a path to do work in an universal and efficient way. 

Fortunately, Golang can support [FP](https://blog.logrocket.com/functional-programming-in-go/#:~:text=The%20goal%20of%20functional%20programming%20is%20to%20make%20the%20state,that%20may%20cause%20side%20effects.)
patterns, and once you common FP functions like [map, find, filter, and reduce](https://medium.com/@geisonfgfg/functional-go-bc116f4c96a4) you'll want to reach for these tools to process your data.  

The native functions may not be enough. Im which case, a library like [lo](https://github.com/samber/lo) is worth importing into your project. This `lo` library also the ability to use FP functions within goroutines so that they extremely quick! It also interesting to look at the code behind the FP function like [find](https://github.com/samber/lo/blob/master/find.go) and see how he used [generics](https://go.dev/blog/intro-generics) to make this happen. The documentation for [Find](https://pkg.go.dev/github.com/samber/lo#Find) is nice too. 

For this project, I'm fine without the `lop` goroutine versions. Since I know FP, I can follow the pattern to create a [filter criteria](https://stackoverflow.com/questions/3230944/what-does-predicate-mean-in-the-context-of-computer-science) but instead of boolean predicate function, I want to return a string.

In the stackoverflow, I have a `find` example... 

```
str, ok := lo.Find([]string{"foobar"}, func(i string) bool {
    return i == "b"
})
```

I want to pass into a list of order status and return an order description by the id. This is the filter criteria that I want to use. 

```
orderStatus, ok := lo.Find(order_status_list, func(orderStatus OrderStatus) bool {
		return orderStatus.Id == FRAUD_CHECKED
	})
```

Instead of struggling to learn `lo` in a larger project, I kicked its tires in a [go playgroound](https://go.dev/play/p/dxXmzkdCRPw).


```
CREATE TABLE customer_journey(
  order_status_id serial primary key,
  order_status varchar(255) NOT NULL,
  customer_id varchar(255), 
  order_id INT,
  CONSTRAINT customer_id_fkey FOREIGN KEY (customer_id)
    REFERENCES customers (customer_id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT order_id_fkey FOREIGN KEY (order_id)
    REFERENCES orders (order_id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);
```

## Day 4

I wanted to continue follow the next tutorial Aurelie Vache series and setup the CLI so that I could play with my new API and simulate customers placing orders and having the shipping respond with fake status. 

I ran into a problem installing the recommend [Cobra](https://github.com/spf13/cobra) library which exposed that my $GOPATH wasn't setup properly. 

I explored and fixed my $GOPATH like so:

which go
/Users/braflow/.gvm/gos/go1.19.1/bin/go

export GOPATH=/Users/braflow/.gvm/gos/go1.19.1/bin/go
export PATH=$PATH:$GOPATH/bin

nano .zshrc 

At the bottom of the file, I added

```
export GOPATH=/Users/braflow/.gvm/gos/go1.19.1/bin/go
export PATH=$PATH:$GOPATH/bin
```

Then save & exit nano via: Ctrl o & Ctrl x

Now when I try 

echo $GOPATH 

I should see `/Users/braflow/.gvm/gos/go1.19.1/bin/go`

If not, you may to restart the terminal - obviously your user path will be different than mine.

And my problem still exists! I'm reading about GOPATH vs GOROOT and seeing that this approach seems to deprecated!

```
go install github.com/spf13/cobra-cli@latest
go: could not create module cache: stat /Users/braflow/.gvm/gos/go1.19.1/bin/go/pkg/mod: not a directory
```
I know mod is a directory that exists with read/write access so what gives?!

I was able to find cobra-cli file within my file system so I simply manually copied to over the `bin` folder within my `cli` directory

```
northwind-frostpunk/cli on ??? cobra [!?] via ???? v1.19.1
??? ./bin/cobra-cli
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.
```
So it seems to run but the init fails 

```
??? ./bin/cobra-cli init
Error: exit status 1
```

well... it might also be because I'm on the latest version of go according to this [github issue thread](https://github.com/spf13/cobra/issues/1587); at leaset I'm not the only one.

So I installed 1.18.3
```
gvm install go1.18.3
gvm use go1.18.3
```
And returned to my folder and presto! 
```
cd northwind-frostpunk/cli
go install github.com/spf13/cobra-cli@latest
cobra-cli init
Your Cobra application is ready at
/Users/braflow/northwind-frostpunk/cli
```
I'm back on track with Vache's tutorial except instead of her `get` command I want to create a command that will move the game to the next day

```
cobra-cli add advanceDay
```

Perhaps, I'll want to rewind the day but for now let's MVP and move forward. 

```
go run main.go advanceDay
```

Ok... so now I have 2 terminal windows.

In the first, I'm running the API so:

```
cd northwind-frostpunk
go run .
```

and in the second, I want to test the API using the CLI

```
cd northwind-frostpunk/cli
go run main.go advanceDay
```

I discovered I needed to do a bit more work to [present the json](https://stackoverflow.com/questions/20873640/how-to-get-json-object-by-calling-a-url-in-go-language) 

```
northwind-frostpunk/cli on ??? cobra [!] via ???? v1.19.1 
??? go run . advanceDay day3
This is the argument that I passed: day3
Perfect! I should see products:  map[body:[map[product_name:Chai unit_price:18 units_in_stock:39] map[product_name:Chang unit_price:19 units_in_stock:17] map[product_name:Aniseed Syrup unit_price:10 units_in_stock:13]] description: status:200]
```
# Day 5

https://github.com/inancgumus/learngo

planning 5 days of orders to reach 1M+ in sales. Is this even possible with the current customer base?

And what is already in the database?

```
SELECT COUNT(*)
FROM customers;
```
91 Customers
```
SELECT COUNT(*)
FROM products;
```
77 Products
```
SELECT COUNT(*)
FROM orders;
```
830 Orders

Let's look at few orders
```
SELECT * 
FROM orders
LIMIT 5;
```
Can I copy and paste the pgAdmin4 table and convert it to markdown? Yes [tableconvert](https://tableconvert.com/) is excellent. 

```
| "order_id" | "customer_id" | "employee_id" | "order_date" | "required_date" | "shipped_date" | "ship_via" | "freight" | "ship_name"                 | "ship_address"         | "ship_city"      | "ship_region" | "ship_postal_code" | "ship_country" |
|------------|---------------|---------------|--------------|-----------------|----------------|------------|-----------|-----------------------------|------------------------|------------------|---------------|--------------------|----------------|
| 10248      | "VINET"       | 5             | "1996-07-04" | "1996-08-01"    | "1996-07-16"   | 3          | 32.38     | "Vins et alcools Chevalier" | "59 rue de l'Abbaye"   | "Reims"          |               | "51100"            | "France"       |
| 10249      | "TOMSP"       | 6             | "1996-07-05" | "1996-08-16"    | "1996-07-10"   | 1          | 11.61     | "Toms Spezialit??ten"        | "Luisenstr. 48"        | "M??nster"        |               | "44087"            | "Germany"      |
| 10250      | "HANAR"       | 4             | "1996-07-08" | "1996-08-05"    | "1996-07-12"   | 2          | 65.83     | "Hanari Carnes"             | "Rua do Pa??o, 67"      | "Rio de Janeiro" | "RJ"          | "05454-876"        | "Brazil"       |
| 10251      | "VICTE"       | 3             | "1996-07-08" | "1996-08-05"    | "1996-07-15"   | 1          | 41.34     | "Victuailles en stock"      | "2, rue du Commerce"   | "Lyon"           |               | "69004"            | "France"       |
| 10252      | "SUPRD"       | 4             | "1996-07-09" | "1996-08-06"    | "1996-07-11"   | 2          | 51.3      | "Supr??mes d??lices"          | "Boulevard Tirou, 255" | "Charleroi"      |               | "B-6000"           | "Belgium"      |
```
I like the third row featuring the Brazilian customer who has $65.83 in freight charges. I wonder what they bought?

The customer id is HANAR and the order id is 10250

What is the customer's name?

```
SELECT C.customer_id, C.contact_name, C.city
FROM customers AS C
WHERE customer_id = 'HANAR';
```

Mario Pontes, Accounting Manager from Rio de Janeiro

What did this customer buy?

```
SELECT OD.order_id, OD.product_id, OD.quantity 
FROM order_details AS OD
WHERE order_id = 10250;
```

```
SELECT O.order_id, O.order_date, O.shipped_date, O.freight, O.ship_name 
FROM orders AS O
WHERE order_id = 10250;

```

Order 10250 was ordered on 1996-07-08 and shipped on 1996-07-12 via Hanari Carnes for $65.83 in freight charges. 

I will need to write a join to get the product names

```
SELECT C.contact_name 
      ,ODE.order_id
	    ,PRO.product_name
	    ,ODE.quantity
	    ,ODE.unit_price
  FROM order_details AS ODE
  JOIN orders AS ORD
    ON ORD.order_id = ODE.order_id
  JOIN customers AS C
    ON ORD.customer_id = C.customer_id
  JOIN products AS PRO
    ON PRO.product_id = ODE.product_id
 WHERE ODE.order_id = 10250
 ORDER BY ODE.order_id ASC; 
```

Now, I need to use this query as a sub query so that I create a final total for the products.

This is a simple way to start considering sub queries. We can repeat the above result set but as sub query.

```
SELECT * 
FROM (SELECT C.contact_name 
       ,ODE.order_id
	   ,PRO.product_name
	   ,ODE.quantity
	   ,ODE.unit_price
  FROM order_details AS ODE
  JOIN orders AS ORD
    ON ORD.order_id = ODE.order_id
  JOIN customers AS C
    ON ORD.customer_id = C.customer_id
  JOIN products AS PRO
    ON PRO.product_id = ODE.product_id
 WHERE ODE.order_id = 10250
 ORDER BY ODE.order_id ASC) as CO;
```

You will see the CO (our new CustomerOrder table) has the same result as the sub query. But I want to aggregate the data and sum the total of unit price multiplied against the quantity. 

SELECT (CO.customer as name) 
FROM (SELECT C.contact_name 
       ,ODE.order_id
	   ,PRO.product_name
	   ,ODE.quantity
	   ,ODE.unit_price
  FROM order_details AS ODE
  JOIN orders AS ORD
    ON ORD.order_id = ODE.order_id
  JOIN customers AS C
    ON ORD.customer_id = C.customer_id
  JOIN products AS PRO
    ON PRO.product_id = ODE.product_id
 WHERE ODE.order_id = 10250
 ORDER BY ODE.order_id ASC) as CO;


Can I join the customer and order tables to produce a record?

https://www.sisense.com/blog/how-to-format-numbers-as-currency-in-postgres-mysql-and-redshift/
https://database.guide/how-to-format-numbers-as-currency-in-postgresql/

cannot cast type real to money!

```
  cast(to_char(sum(CO.unit_price * CO.quantity),'L99D99') as money) as total
```

ERROR:  invalid input syntax for type money: "$ ##.##"

If I remove the quantity I get $66.90 with only the unit price. 

I had the wrong format! Since my total will be over $100, I needed a larger number format like 'L999G999D99'

```
SELECT CO.contact_name AS name,
     string_agg(CO.product_name, ', ') AS product_list,
	   cast(to_char(sum(CO.unit_price * CO.quantity),'L999G999D99') as money) as total
FROM (SELECT C.contact_name 
       ,ODE.order_id
	   ,PRO.product_name
	   ,ODE.quantity
	   ,ODE.unit_price
  FROM order_details AS ODE
  JOIN orders AS ORD
    ON ORD.order_id = ODE.order_id
  JOIN customers AS C
    ON ORD.customer_id = C.customer_id
  JOIN products AS PRO
    ON PRO.product_id = ODE.product_id
 WHERE ODE.order_id = 10250
 ORDER BY ODE.order_id ASC) as CO
 GROUP BY CO.contact_name; 
````
Now I see the desired result:
```
Mario Pontes purchased a $1,813.00 of products including Jack's New England Clam Chowder, Manjimup Dried Apples, and Louisiana Fiery Hot Pepper Sauce.
``
Stretch goal: instead of sub query, could I could the same thing with a [CTE](https://learnsql.com/blog/sql-subquery-cte-difference/)?

That is easy enough. All I had to do was reposition the sub query.

```
/*
creates a table with one customer and shows the products they ordered.
The table will feature one product per row
*/

WITH CO AS (
	SELECT C.contact_name 
		   ,ODE.order_id
		   ,PRO.product_name
		   ,ODE.quantity
		   ,ODE.unit_price
	  FROM order_details AS ODE
	  JOIN orders AS ORD
		ON ORD.order_id = ODE.order_id
	  JOIN customers AS C
		ON ORD.customer_id = C.customer_id
	  JOIN products AS PRO
		ON PRO.product_id = ODE.product_id
	 WHERE ODE.order_id = 10250
	 ORDER BY ODE.order_id ASC)

-- review the customer's total order and their product list in one row 

SELECT CO.contact_name AS name,
	   string_agg(CO.product_name, ', ') AS product_list,
	   cast(to_char(sum(CO.unit_price * CO.quantity),'L999G999D99') as money) as total
FROM CO
 GROUP BY CO.contact_name;  
```

I like this CTE approach as it feels more natural as to how I built it. I started with the query to reduce my table then used to the second to query to reduce it further. The comments are also useful to organize and document the SQL. 

The version 1.18 of [Golang introduces generics](https://itnext.io/generic-map-filter-and-reduce-in-go-3845781a591c) and several functions - `golang.org/x/exp/slices` - found in the functional programming so I may not need a helper library like [lo](https://github.com/samber/lo).

```
import (
  "golang.org/x/exp/slices"
)

// we can use slices to search through an array and find the struct by its key
idx := slices.IndexFunc(countCustomerRows, func(c CountCustomerRow) bool { return c.Key == "key1" })
```

I think I prefer the API of `lo` though, and prepared a [go playground example for it](https://github.com/samber/lo/issues/226) and enjoyed collaborating with Sambre to make it better!

# Day 6

I'll regain focus and return to interacting with the API via the CLI

I can run the advanceDay command like so: 

```
go run . advanceDay
```

But really I should start a command that's starts the experience so I've developed this:

```
go run . startGame 
```

With this endpoint, I can start writing the game but will leave that for another day. 

# Day 7

As I grow more confident with my new sub query SQL skills, I can start to apply it to real world problems. I wanted to write a query in BigQuery that would mine our daily Java logs. 

First, I setup a data sink from Logs Explorer to BigQuery which would collect the logs for successful and failed orders. In BigQuery, I had three query tabs open. Before writing this final query, I wrote the success and fail queries in separate tabs, and then combined them in the third tab.  

I'll share the query removing any sensitive content: 

```
WITH 
  success_revenue AS (
  SELECT
   '1' AS id,
    CAST(
      SUM(ROUND(CAST(REGEXP_SUBSTR(jsonPayload.message, "success-find-money-pattern-here(.*?[^]]*)") AS FLOAT64))) AS STRING FORMAT '$999,999.99'
    ) AS success_total,
    COUNT(ROUND(CAST(REGEXP_SUBSTR(jsonPayload.message, "success-find-money-pattern-here(.*?[^]]*)") AS FLOAT64))) AS success_count
  FROM
    `DATA-SINK-TABLE-HERE`
  WHERE (jsonPayload.message LIKE '%success-string-pattern-here%')
  ),

  failed_revenue AS (
  SELECT
   '1' AS id,
      CAST(
        SUM(ROUND(CAST(REGEXP_SUBSTR(jsonPayload.message, "fail-find-money-pattern-here(.*?[^]]*)") AS FLOAT64),2)) AS STRING FORMAT '$999,999.99'
      ) AS Failed_Total,
    COUNT(ROUND(CAST(REGEXP_SUBSTR(jsonPayload.message, "fail-find-money-pattern-here(.*?[^]]*)") AS FLOAT64),2)) AS failed_Count
  FROM
    `DATA-SINK-TABLE-HERE`
  WHERE NOT (jsonPayload.message LIKE '%success-string-pattern-here%')),

  all_revenue
    as
    (
      select  success_revenue.id,  
              success_revenue.success_total,
              success_revenue.success_count,
              failed_revenue.failed_total,
              failed_revenue.failed_count
      from success_revenue
      JOIN failed_revenue ON success_revenue.id=failed_revenue.id
    )

Select * from all_revenue
```
I used a regex to pull the totals from each string and where clause to separate the success from the failed orders.  I also counted the strings to get a sense of frequency. This query produced a single row with both success and failed totals and counts. 

I had some help from these blog posts: [write better CTE](https://towardsdatascience.com/common-table-expressions-5-tips-for-data-scientists-to-write-better-sql-bf3547dcde3e), [multiple cte](https://learnsql.com/blog/how-to-use-two-ctes-in-sql/), and [BiqQuery CTE](https://popsql.com/learn-sql/bigquery/how-to-write-a-common-table-expression-in-bigquery). How cool is [popsql](https://popsql.com/)?!

Similiar to importing class, I would hope its possible to break CTEs out into their own files for re-use and to isolate their value, and then import that work without having to maintain one large file but I don't see that [documented](https://www.postgresql.org/docs/13/queries-with.html) nor any posts about it so far. I'll continue the search while tackling [Recursive CTE](https://www.fusionbox.com/blog/detail/graph-algorithms-in-a-database-recursive-ctes-and-topological-sort-with-postgres/620/) [Graph Queries](https://medium.com/white-prompt-blog/implementing-graph-queries-in-a-relational-database-7842b8075ca8) next.

# Day 8 

In order to get to Graph theory, I need to start the graph - start the game! The user begins at the first node in the graph, and then will make decisions to grow it so that I can traverse it with SQL and see what they did along the way.

The [Interactive](https://dev.to/divrhino/building-an-interactive-cli-app-with-go-cobra-promptui-346n) [CLI](https://www.thorsten-hans.com/lets-build-a-cli-in-go-with-cobra/) needs to provide options so the user can simply use their arrows keys to select their choice.


So northwind-frostpunk is a bit much to type of each to run the CLI. I want something shorter like battery since I've just decided this is a game about the birth of lithium mining in Canada. Our miners will be using our Northwind database to purchase supplies.  

Once I updated the battery command, then I can build the project 

```
go build . 
```
This produced a battery.go file in my pkg > battery folder but I can't use the `battery` command just yet...

I want to be able to type:

```
battery startGame
```

which doesn't work yet so I'll use:

```
go run . startGame 
```

# Day 9

By now, I have the API and CLI working together. I can issue commands and interact with the API, and I'm using the CLI to drive. 

This has many been a solo effort but I'm starting to think about security and consequences of making it public. The end goal would be share this story with the world; the beauty (and occasional horror) of any internet-enabled app. 

<img src="ai_co_author.png" />

fmt.Println("Before sleep the time is:", time.Now().Unix())     // Before sleep the time is: 1257894000
    time.Sleep(2 * time.Second)                                     // pauses execution for 2 seconds
    fmt.Println("After sleep the time is:", time.Now().Unix())  

[sleep in golang](https://golangdocs.com/sleep-function-in-golang)

## Links

- [quick security wins](https://www.rapid7.com/blog/post/2016/07/13/quick-security-wins-in-golang/)

## Generics - 

I want to leverage [generics](https://www.digitalocean.com/community/tutorials/how-to-use-generics-in-go) to pass in either string or int for user input. The user is supposed to enter a int but what happens they enter a string?!


```
func validateOption[T any](v T) bool {
```

```
func validateOption[T comparable](v T) bool {
```

I used the [reflect](https://freshman.tech/snippets/go/check-type-of-value/) package to determine that all input is a string whether the user types 6 or apple.


valid json 
```
{"Name": "Ed"}
```
This is a JSON stream not valid json!
```
{"Name": "Ed"}{"Name": "Sam"}{"Name": "Bob"}
```
valid json
```
[{"Name": "Ed"},{"Name": "Sam"},{"Name": "Bob"}]
```

I need to create response which converts my query results into an array of product structs. I'm learning all about [marshalling](https://socketloop.com/tutorials/golang-unmarshal-json-from-http-response) as well as the [frustration] (https://bluehive.medium.com/json-cannot-unmarshal-object-into-a-go-struct-field-11fadb1a2a94) and [pitfalls](https://ahmet.im/blog/golang-json-decoder-pitfalls/ )while working with go.

I get the following error on my initial implementation: 

```
panic: json: cannot unmarshal object into Go struct field ProductsHttpResp.body of type []battery.Product
```

After reviewing this blog post about mismatching types, I realized that I had used the wrong get function to unmarshal my response. Instead of `getUrl`, I needed to use the corresponding and more specific `GetUrlProductsResponse`. In the future, I'll refactor both of these into the a common `getUrl` which could return multiple response structs through Generics. 

https://mj-go.in/golang/async-http-requests-in-go#synchronous-http-requests

Where there are copious amounts of posts on writing concurrent go with goroutines, I found it challenging to find ones on writing go http calls in series. I wanted to send an instruction to the server, and based on that instruction make another api call. 

First, I tried to search for series or sequential or synchronous or blocking request and finally found what I was looking based on a "chain" search tieing to back my mental of Javascript and how we can chain together promises. 

Trust in your mental model! I was confident that I could make multiple API calls in series similar to Java and Python.

When I attempted, I couldn't seem to process the response correctly. After a few hours of wrestling with issue chasing many red herrings, I finally compared the json response in a browser from two tabs.

```
{
status: 200,
description: "Hey headwinds",
body: [
        {
        product_name: "Chai",
        unit_price: 18,
        units_in_stock: 39
        },
```
I double nested the body!
```
{
status: 200,
description: "",
body: {
  status: 200,
  description: "Hey headwinds",
  body: [
          {
          product_name: "Chai",
          unit_price: 18,
          units_in_stock: 39
          },
```
No wonder when I attempted to unmarshall it got confused since it didn't match expect response struct. 

Since I'm learning languages these hurdles to encourage absorption; the puzzle makes me want to chase after a solution and here I am at 10:30pm after starting at 8:30am still excited about the final solution. 

# Day 10

## Testing

After 2 weeks of exloring golang, I think it's time to circle back to test instead pursuing the remainder of the build plan; in fact testing will always be part of the early build but I wanted to learn more the language before investing in tests. I'm ready. As a prerequisite, this [blog post How To Write Unit Tests in Go](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package
) recommends "A familiarity with the Go programming language" before writing your first unit test. 

Similar to Python, Java, or Javascript, I want to build tests that produce reports and give me a sense of code coverage. I'm not aiming for 100% coverage but want to test most important flows that make up the core experience. With tests in place, I have confidence to make system tweaks. The lack of tests already bit me yesterday I updated a route path and suddenly my product serivce failed. 

So far, all the tests I've written attempt to produce a single result like [mocking a request](https://medium.com/zus-health/mocking-outbound-http-requests-in-go-youre-probably-doing-it-wrong-60373a38d2aa). I wonder if I could use unit testings strategies more like A/B testing though? Would there be any value in writing tests that demonstrate several different results such as testing ways to [optimize SQL queries](https://webapp.io/blog/postgres-query-speedups/)?

### Links

- [go lang httptest](https://golang.cafe/blog/golang-httptest-example.html)
- [unit testing](https://medium.com/@victorsteven/understanding-unit-and-integrationtesting-in-golang-ba60becb778d)
- [function declaration](https://stackoverflow.com/questions/34031801/function-declaration-syntax-things-in-parenthesis-before-function-name)
- [tests for database](https://markphelps.me/posts/writing-tests-for-your-database-code-in-go/)
- [go style guide](https://google.github.io/styleguide/go/guide)
- [go chaining http handlers](https://medium.com/@goncharovny/how-to-chain-http-handlers-in-go-33c96396b397)

# Day 11

## Testing vs Pushing a prototype?!

After a day of testing Go controllers and not making as much as progress as I would have liked, I've decided to switch back to building out the services first, and try to unit test small functions and save the controllers for later. 

I want to see how far my user testers get before I spend too much time unit testing flows, and adopt the approach of adding tests where it breaks.   

The more work within the terminal, the more I want to bring design into this space to format the text; make it pleasing to the eye taking inspiration from tools like [charm](https://charm.sh/)

I was able to copy and paste the sql from the northwind-psql project into the elephantsql browser and after executing it, it creates the Northwind tables and seeds them.

Similar to Python or Java, I want to use `.env` file which isn't committed to git to store connection strings and other [variables](https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66). It's tempting to use godotenv but with go with Viper since I'm already using it in the CLI

I have 2 hosting options that I'm most familar with: either Vercel or GCP. I could either host this Go API as a Cloud Run container on GCP or I could [deploy to Vercel](https://vercel.com/docs/concepts/functions/serverless-functions/runtimes/go).

Now that I'm feeling more confidedent writing Go, and already have one Go app (a feed file system reporting tool) under my belt at [Loblaw Digital](https://www.loblawdigital.co/), I'm curious about the market may hold for Go developers worldwide. There certaining some interesting positions out there like from this one from [semitechnologies](https://semitechnologies.teamtailor.com/jobs/1898110-software-engineer-golang) which has a vector search tool using NLP. As fascinating it is to look at sites like [levels.fyi](https://www.levels.fyi/) and share with my manager (we could use banding as there really is no advancement beyond senior), the goal is to continue to build up the culture though. 

If you at times you feel bored in your current role, augmenting you daily routine with passion projects like this can smooth those lulls. My morning rituat is have coffee and crush on projecs like this to prepare me for the day of wrestling with eCommerce puzzles; it's a beautiful balanace; it's living the dream of writing both sides of the contract. 
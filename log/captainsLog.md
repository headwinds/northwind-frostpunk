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

#### Log 🚀

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

#### Log 🚀 

-- Sept 22/2022 -- 

I had less time today than yesterday with work heating up. At a large enterprise, at times we only have 60% of a solution ready to serve our trading team and call centre. Developers often need to refer to SOP (service operationg procedures) that document how to run scripts in our admin tool to patch problems where a bit of UI and a button could have saved the day to complete the feature. 

As someone who cares about product and fullstack development, there is where I see large gaps in our UI and the services we offer.  Obviously, I'm proud of my own SOP contributions. 

This write up was banged out to [Nina Las Vegas b2b Swick DJ Set | Keep Hush Live](https://www.youtube.com/watch?v=8d0mHipC1LA) and [HAAi B2B Daniel Avery | Boiler Room x Glitch Festival 2022](https://www.youtube.com/watch?v=WFexmmlpXe0).

## Day 3 

goal 1 min CLI gameplay! 

coming soon 

Points gamification & order status

[poslog](https://www.yellowfinbi.com/blog/2013/06/yfcommunitynews-poslog-one-data-container-to-rule-them-all-well-at-least-for-retail-141038/)

[Composition](https://medium.com/bitaksi-tech/object-oriented-approach-in-go-179c85486c76) 

```
If a struct, field, or function’s name starts with a lowercase letter, that means it’s private
```

to GCP?

https://medium.com/@crypto-gopher/learn-how-to-containerize-your-go-app-in-5-minutes-e654fdb2afd8

#### Order Status

```
CREATE TABLE order_status(
  order_status_id serial primary key,
  CONSTRAINT order_id_fkey FOREIGN KEY (order_id)
    REFERENCES employee (order_id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION,
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
{"status":200,"description":"","body":[{"order_id":"10248","customer_id":"VINET","ship_city":"Reims"},{"order_id":"10249","customer_id":"TOMSP","ship_city":"Münster"},{"order_id":"10250","customer_id":"HANAR","ship_city":"Rio de Janeiro"}]}
```

<img src="orders_response.png" />
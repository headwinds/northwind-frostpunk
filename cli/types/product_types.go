package types

/*
[
{
product_name: "Queso Cabrales",
unit_price: 21,
units_in_stock: 22
},
*/

type Product struct{
    ProductName		string  `json:"product_name"`
    UnitPrice		float64 `json:"unit_price"`
    UnitsInStock	int  	`json:"units_in_stock"`
}
type ProductsHttpResp struct{
    Status      int         `json:"status"`
    Description string      `json:"description"`
    Body        []Product 	`json:"body"`
}
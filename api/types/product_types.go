package types

/*
product_name: "Aniseed Syrup",
unit_price: 10,
units_in_stock: 13
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

type HttpResp struct{
    Status      int         `json:"status"`
    Description string      `json:"description"`
    Body        interface{} `json:"body"`
}
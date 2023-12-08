package types

/*
product_name: "Aniseed Syrup",
unit_price: 10,
units_in_stock: 13
*/

type ProductJson struct{
    ProductName		string  `json:"product_name"`
    UnitPrice		float64 `json:"unit_price"`
    UnitsInStock	int  	`json:"units_in_stock"`
}

type ProductsBody []Product

type ProductsHttpResp struct{
    Status      int         `json:"status"`
    Description string      `json:"description"`
    Body        ProductsBody 	`json:"body"`
}

type HttpResp struct{
    Status      int         `json:"status"`
    Description string      `json:"description"`
    Body        interface{} `json:"body"`
}

type Product struct {
  ProductId       int     `db:"product_id"`
  ProductName     string  `db:"product_name"`
  SupplierId      int     `db:"supplier_id"`
  CategoryId      int     `db:"category_id"`
  QuantityPerUnit string  `db:"quantity_per_unit"`
  UnitPrice       float64 `db:"unit_price"`
  UnitsInStock    int     `db:"units_in_stock"`
  UnitsonOrder    int     `db:"units_on_order"`
  ReorderLevel    int     `db:"reorder_level"`
  Discontinued    int     `db:"discontinued"`
}

type JsonDataResponse struct {
  Type    string    `json:"type"`
  Message string    `json:"message"`
  Data    []Product `json:"data"`
}
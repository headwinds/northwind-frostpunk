package customer_journey

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpResp struct{
    Status      int         `json:"status"`
    Description string      `json:"description"`
    Body        interface{} `json:"body"`
}

type Customer struct{
    CustomerId	string  `json:"customer_id"`
}

type CountCustomerRow struct{
    Count int
}

type TotalCustomersResponseBody struct{
	Total int
}

type DatabaseHandler struct {
	db *sql.DB
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CustomerJourneyController(db *sql.DB) *DatabaseHandler {
	return &DatabaseHandler{
		db: db,
	}
}

func (h *DatabaseHandler) GetTotalCustomers(w http.ResponseWriter, r *http.Request){
	if err := h.db.Ping(); err != nil {
		fmt.Println("DB Error")
	}

	var countCustomerRows []CountCustomerRow
	rows, err := h.db.Query(`SELECT "customer_id" FROM "customer_journey" LIMIT 3`)
	checkErr(err)

	defer rows.Close()
	for rows.Next() {

		var countCustomerRow CountCustomerRow
		err = rows.Scan(&countCustomerRow.Count)
		checkErr(err)
	
		countCustomerRows = append(countCustomerRows, countCustomerRow)
	}

	
	//fmt.Println("customer journey message: ", message)
	idx := slices.IndexFunc(countCustomerRows, func(c CountCustomerRow) bool { return c.Key == "key1" })


	responseBody := TotalCustomersResponseBody{ Total: CountCustomerRow, Message: message }

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(HttpResp{Status: 200, Body: responseBody })

	checkErr(err)

}

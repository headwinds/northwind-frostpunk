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

type CustomerJourney struct{
    CustomerId	string  `json:"customer_id"`
}

type ResponseBody struct{
    CustomJournies []CustomerJourney
	Message string
	FilteredOrderStatusList []OrderStatus
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

func (h *DatabaseHandler) GetCustomerJournies(w http.ResponseWriter, r *http.Request){
	if err := h.db.Ping(); err != nil {
		fmt.Println("DB Error")
	}

	var customJournies []CustomerJourney
	rows, err := h.db.Query(`SELECT "customer_id" FROM "customer_journey" LIMIT 3`)
	checkErr(err)

	defer rows.Close()
	for rows.Next() {

		var customerJourney CustomerJourney
		err = rows.Scan(&customerJourney.CustomerId)
		checkErr(err)
	
		customJournies = append(customJournies, customerJourney)
	}

	message := GetStatusDescriptions(ACK_RECEIVED_FROM_3PL)
	fmt.Println("customer journey message: ", message)

	responseBody := ResponseBody{ CustomJournies: customJournies, Message: message }

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(HttpResp{Status: 200, Body: responseBody })

	checkErr(err)

}

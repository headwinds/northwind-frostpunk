package customer_journey

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/headwinds/northwind-frostpunk/api/controllers/orders"
	"golang.org/x/exp/slices"
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
	FilteredOrderStatusList []orders.OrderStatus
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

func getStatusDescriptions(id orders.Status) string {
	orderStatusList := orders.GetOrderStatusList()

	// finds the index in the list of the order status we are interested in
	orderStatusIdx := slices.IndexFunc(orderStatusList, func(orderStatus orders.OrderStatus) bool { return orderStatus.Id == id })
	orderStatus := orderStatusList[orderStatusIdx]

	return orderStatus.Description;
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

	message := getStatusDescriptions(orders.ACK_RECEIVED_FROM_3PL)
	fmt.Println("customer journey message: ", message)

	responseBody := ResponseBody{ CustomJournies: customJournies, Message: message }

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(HttpResp{Status: 200, Body: responseBody })

	checkErr(err)

}

package game

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/headwinds/northwind-frostpunk/api/utils"
)

func NextTurnController(db *sql.DB) *DatabaseHandler {
	return &DatabaseHandler{
		db: db,
	}
}

func (h *DatabaseHandler) NextTurn(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Ping(); err != nil {
		fmt.Println("DB Error")
	}

	decision := r.URL.Query().Get("decision")
	fmt.Println("decision =>", decision) // I should consider encoding each decision into an id


	//products
	url := "http://localhost:8080/products/view?page=1&limit=10"
	jsonData := utils.GetUrlProductsResponse(url)

	if jsonData.Status == 200 {
		// we don't want to double nest the response body so we discard the status and description
		// so we could do
		// //json.NewEncoder(w).Encode(types.HttpResp{Status: 200, Body: jsonData.Body})
		// or simply do json.NewEncoder(w).Encode(jsonData)

		w.Header().Set("Content-Type", "application/json")
		
		json.NewEncoder(w).Encode(jsonData)
    }

	
	

}

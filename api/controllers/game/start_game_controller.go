package game

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

type DatabaseHandler struct {
	db *sql.DB
}

func StartGameController(db *sql.DB) *DatabaseHandler {
	return &DatabaseHandler{
		db: db,
	}
}

func (h *DatabaseHandler) StartGame(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Ping(); err != nil {
		fmt.Println("DB Error")
	}

	w.Header().Set("Content-Type", "application/json")
	gameDay := GameDayManager()
    json.NewEncoder(w).Encode(HttpResp{Status: 200, Body: gameDay})
}
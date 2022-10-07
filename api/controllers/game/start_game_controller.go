package game

import (
	"database/sql"
	"fmt"
	"net/http"
)

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

	w.Write([]byte("Hello, World"))
}
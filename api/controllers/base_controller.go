/*
this is an excellent pattern and I have refactored my controllers based on it
https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang
*/

package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
)

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	db *sql.DB
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

// HelloWorld returns Hello, World
func (h *BaseHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Ping(); err != nil {
		fmt.Println("DB Error")
	}

	w.Write([]byte("Hello, World"))
}
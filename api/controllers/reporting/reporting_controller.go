package reporting

import (
	"database/sql"
	"fmt"
	"net/http"
)

// BaseHandler will hold everything that controller needs
type DatabaseHandler struct {
	db *sql.DB
}

// NewBaseHandler returns a new BaseHandler
func ReportingController(db *sql.DB) *DatabaseHandler {
	return &DatabaseHandler{
		db: db,
	}
}

// Report will handle the preparing the data to be presented
func (h *DatabaseHandler) Report(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Ping(); err != nil {
		fmt.Println("DB Error")
	}

	w.Write([]byte("Hello, World"))
}
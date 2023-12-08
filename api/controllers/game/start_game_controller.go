package game

import (
	"encoding/json"
	"fmt"
	"github.com/headwinds/northwind-frostpunk/api/types"
	"net/http"

  "github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseHandler struct {
	dbpool *pgxpool.Pool
}

func StartGameController(dbpool *pgxpool.Pool) *DatabaseHandler {
	return &DatabaseHandler{
    dbpool: dbpool,
	}
}

func (h *DatabaseHandler) StartGame(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Hit /game/start route")
	gameDay := GameDayManager()

	json.NewEncoder(w).Encode(types.HttpResp{Status: 200, Body: gameDay})
}

package types

import (
  "github.com/jackc/pgx/v5/pgxpool"
)

type JsonMessageResponse struct {
  Type    string `json:"type"`
  Message string `json:"message"`
}

type DatabaseHandler struct {
  Connpool *pgxpool.Pool
}
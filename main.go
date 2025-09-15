package main

import (
	"database/sql"
	"log"

	"github.com/go-api-payline/api"
	db "github.com/go-api-payline/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:secret@localhost:5432/payline_v1?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db ", err)
	}

queries := db.New(conn)
server := api.NewServer(queries)


err = server.Start(serverAddress)
if err != nil {
	log.Fatal("cannot start server:", err)
}

}

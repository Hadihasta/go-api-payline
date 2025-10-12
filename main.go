package main

import (
	"database/sql"
	"log"

	"github.com/go-api-payline/api"
	db "github.com/go-api-payline/db/sqlc"
	"github.com/go-api-payline/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db ", err)
	}

	queries := db.New(conn)
	server := api.NewServer(queries)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}

package main

import (
	"database/sql"
	"log"

	"github.com/Emmrys-Jay/simple-bank/api"
	db "github.com/Emmrys-Jay/simple-bank/db/sqlc"
	"github.com/Emmrys-Jay/simple-bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAdress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}

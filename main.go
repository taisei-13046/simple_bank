package main

import (
	"database/sql"
	"log"

	"github.com/taisei-13046/simplebank/api"
	db "github.com/taisei-13046/simplebank/db/sqlc"
	"github.com/taisei-13046/simplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("connot start server", err)
	}
}

package app

import (
	"book-keeper/handlers"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var dbClient *sql.DB

func Start() {
	router := mux.NewRouter()
	dbClient = getDbClient()
	defer dbClient.Close()
	rh := handlers.NewReaderHandler(dbClient)
	router.HandleFunc("/readers", rh.GetAllReaders).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func getDbClient() *sql.DB {

	dbURL := "host=localhost dbname=booking_keeper user=synergy password=12345 port=5432"
	client, err := sql.Open("pgx", dbURL)
	if err != nil {
		panic(err)
	}
	err = client.Ping()
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

package app

import (
	"book-keeper/handlers"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Start() {
	router := mux.NewRouter()
	dbClient := getDbClient()
	defer dbClient.Close(context.Background())
	rh := handlers.NewReaderHandler(dbClient)
	router.
		HandleFunc("/readers", rh.GetAllReaders).
		Methods(http.MethodGet)
	router.
		HandleFunc("/readers/{id}", rh.GetReaderById).
		Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func getDbClient() *pgx.Conn {

	dbURL := "host=localhost dbname=booking_keeper user=synergy password=12345 port=5432"
	client, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}

	return client
}

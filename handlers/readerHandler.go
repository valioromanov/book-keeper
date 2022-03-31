package handlers

import (
	"book-keeper/service"
	"net/http"

	"github.com/jackc/pgx/v4"
)

type ReaderHandlers struct {
	service service.ReaderService
}

func (rh *ReaderHandlers) GetAllReaders(w http.ResponseWriter, r *http.Request) {
	readers, err := rh.service.GetAllReaders()

	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, err)
	} else {
		WriteResponse(w, http.StatusAccepted, readers)
	}

}

func NewReaderHandler(dbClient *pgx.Conn) ReaderHandlers {
	return ReaderHandlers{service: service.NewReaderService(dbClient)}
}

package handlers

import (
	"book-keeper/service"
	"database/sql"
	"net/http"
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

func NewReaderHandler(dbClient *sql.DB) ReaderHandlers {
	return ReaderHandlers{service: service.NewReaderService(dbClient)}
}

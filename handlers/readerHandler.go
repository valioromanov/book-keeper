package handlers

import (
	"book-keeper/dto"
	"book-keeper/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

func (rh *ReaderHandlers) GetReaderById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	readers, err := rh.service.GetReaderById(id)

	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, err)
	} else {
		WriteResponse(w, http.StatusAccepted, readers)
	}

}

func (rh *ReaderHandlers) NewReader(w http.ResponseWriter, r *http.Request) {

	var reqNewReader dto.NewReaderRequest

	err := json.NewDecoder(r.Body).Decode(&reqNewReader)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}
}

func NewReaderHandler(dbClient *pgx.Conn) ReaderHandlers {
	return ReaderHandlers{service: service.NewReaderService(dbClient)}
}

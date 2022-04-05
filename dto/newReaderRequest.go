package dto

import (
	"book-keeper/errs"
	"book-keeper/model"
	"time"
)

type NewReaderRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	IdentNo   string `json:"identificationNo"`
}

func (nr *NewReaderRequest) ValidateNewReaderRequest() *errs.AppError {
	if nr.FirstName == "" {
		return errs.NewBadRequestError("First name could not be empty!")
	}
	if nr.LastName == "" {
		return errs.NewBadRequestError("Last name could not be empty!")
	}
	if nr.IdentNo == "" {
		return errs.NewBadRequestError("Identification no could not be empty!")
	}
	return nil
}

func (nr *NewReaderRequest) NewReaderRequestToReader() model.Reader {
	reg := time.Now()
	return model.Reader{
		FirstName: nr.FirstName,
		LastName:  nr.LastName,
		IdentNo:   nr.IdentNo,
		RegDate:   &reg,
		Status:    "0",
	}

}

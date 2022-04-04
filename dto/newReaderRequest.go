package dto

import (
	"book-keeper/model"
	"time"
)

type NewReaderRequest struct {
	FirstName string
	LastName  string
	IdentNo   string
}

func (nr *NewReaderRequest) NewReaderRequestToReader() model.Reader {
	reg := time.Now().Format("2006-01-02 15:04:05")
	regTime, _ := time.Parse(reg, "2006-01-02 15:04:05")
	return model.Reader{
		FirstName: nr.FirstName,
		LastName:  nr.LastName,
		IdentNo:   nr.IdentNo,
		RegDate:   &regTime,
		Status:    "0",
	}

}

package model

import "time"

//Reader represents the preson who is registered as reader in the labriray
type Reader struct {
	Id        int64      `json:"id" db:"id"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	IdentNo   string     `json:"identityNo" db:"identity_no"`
	RegDate   *time.Time `json:"registrationDate" db:"registration_date"`
	Status    string     `db:"status"`
}

type ReaderRepository interface {
	FindAll() ([]Reader, error)
}

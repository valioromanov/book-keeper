package model

//Reader represents the preson who is registered as reader in the labriray
type Reader struct {
	Id        string `json:"id" db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	IdentNo   string `json:"identityNo" db:"identity_no"`
	RegDate   string `json:"registrationDate" db:"registration_date"`
	Status    string `db:"status"`
}

type ReaderRepository interface {
	FindAll() ([]Reader, error)
}

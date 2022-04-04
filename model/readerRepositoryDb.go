package model

import (
	"book-keeper/errs"
	"context"

	"github.com/jackc/pgx/v4"
)

type ReaderRepositoryDb struct {
	client *pgx.Conn
}

//function that make a query for getting all readers
func (r *ReaderRepositoryDb) FindAll() ([]Reader, *errs.AppError) {
	allReadersSql := "select * from reader"
	var reader Reader
	readers := make([]Reader, 0)
	rows, err := r.client.Query(context.Background(), allReadersSql)
	if err != nil {

		return nil, errs.NewUnexceptedError(err.Error())
	}
	// rows.Next automatically calls rows.Close()
	// but to be sure we can close it again
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&reader.Id,
			&reader.FirstName,
			&reader.LastName,
			&reader.IdentNo,
			&reader.RegDate,
			&reader.Status)
		if err != nil {
			return nil, errs.NewUnexceptedError(err.Error())
		}
		readers = append(readers, reader)
	}
	return readers, nil
}

func (r *ReaderRepositoryDb) FindById(id string) (*Reader, *errs.AppError) {
	readerById := "select * from reader where id=$1"
	var reader Reader

	err := r.client.QueryRow(context.Background(), readerById, id).Scan(&reader.Id,
		&reader.FirstName,
		&reader.LastName,
		&reader.IdentNo,
		&reader.RegDate,
		&reader.Status)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.NewNotFoundError("Not found reader with id: " + id)
		}
		return nil, errs.NewUnexceptedError(err.Error())
	}

	return &reader, nil
}

func (r *ReaderRepositoryDb) InsertNewReader(read Reader) (*Reader, *errs.AppError) {
	insertNewReaderQuery := "insert into reader(first_name, last_name, ident_no, registration_date, status) values ($1,$2,$3,$4,$5) returning id"
	tx, err := r.client.Begin(context.Background())
	if err != nil {
		return nil, errs.NewUnexceptedError(err.Error())
	}
	lastInsteredId := 0
	defer tx.Rollback(context.Background())
	err = r.client.QueryRow(context.Background(),
		insertNewReaderQuery,
		read.FirstName,
		read.LastName,
		read.IdentNo,
		read.RegDate,
		read.Status).Scan(&lastInsteredId)

	if err != nil {
		return nil, errs.NewUnexceptedError(err.Error())
	}

	err = tx.Commit(context.Background())

	if err != nil {
		return nil, errs.NewUnexceptedError(err.Error())
	}

	reader := Reader{Id: int64(lastInsteredId)}

	return &reader, nil
}

func NewReaderRepositoryDb(dbClient *pgx.Conn) ReaderRepositoryDb {
	return ReaderRepositoryDb{client: dbClient}
}

package model

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type ReaderRepositoryDb struct {
	client *pgx.Conn
}

func (r ReaderRepositoryDb) FindAll() ([]Reader, error) {
	allReadersSql := "select * from reader"
	var reader Reader
	readers := make([]Reader, 0)
	rows, err := r.client.Query(context.Background(), allReadersSql)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		readers = append(readers, reader)
	}
	return readers, nil
}

func NewReaderRepositoryDb(dbClient *pgx.Conn) ReaderRepositoryDb {
	return ReaderRepositoryDb{client: dbClient}
}

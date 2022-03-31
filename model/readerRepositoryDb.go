package model

import (
	"database/sql"
)

type ReaderRepositoryDb struct {
	client *sql.DB
}

func (r ReaderRepositoryDb) FindAll() ([]Reader, error) {
	allReadersSql := "select * from reader"
	var reader Reader
	readers := make([]Reader, 0)
	rows, err := r.client.Query(allReadersSql)
	if err != nil {
		return nil, err
	}
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

func NewReaderRepositoryDb(dbClient *sql.DB) ReaderRepositoryDb {
	return ReaderRepositoryDb{client: dbClient}
}

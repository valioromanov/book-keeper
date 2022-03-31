package service

import (
	"book-keeper/model"
	"database/sql"
)

type ReaderService struct {
	repo model.ReaderRepositoryDb
}

type ReaderServiceInterface interface {
	GetAllReaders() ([]model.Reader, error)
}

func (rs ReaderService) GetAllReaders() ([]model.Reader, error) {
	return rs.repo.FindAll()
}

func NewReaderService(dbClient *sql.DB) ReaderService {
	return ReaderService{model.NewReaderRepositoryDb(dbClient)}
}

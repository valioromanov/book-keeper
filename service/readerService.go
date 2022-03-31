package service

import (
	"book-keeper/model"

	"github.com/jackc/pgx/v4"
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

func NewReaderService(dbClient *pgx.Conn) ReaderService {
	return ReaderService{model.NewReaderRepositoryDb(dbClient)}
}

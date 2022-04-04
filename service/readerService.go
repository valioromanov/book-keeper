package service

import (
	"book-keeper/dto"
	"book-keeper/errs"
	"book-keeper/model"

	"github.com/jackc/pgx/v4"
)

type ReaderService struct {
	repo model.ReaderRepositoryDb
}

type ReaderServiceInterface interface {
	GetAllReaders() ([]model.Reader, *errs.AppError)
	GetReaderById(id string) (*model.Reader, *errs.AppError)
	RegisterReader(nr dto.NewReaderRequest) (*dto.NewReaderResponse, *errs.AppError)
}

func (rs *ReaderService) GetAllReaders() ([]model.Reader, *errs.AppError) {
	return rs.repo.FindAll()
}

func (rs *ReaderService) GetReaderById(id string) (*model.Reader, *errs.AppError) {
	return rs.repo.FindById(id)
}

func (rs *ReaderService) RegisterReader(nr dto.NewReaderRequest) (*dto.NewReaderResponse, *errs.AppError) {
	readerReq := nr.NewReaderRequestToReader()

	dbResp, err := rs.repo.InsertNewReader(readerReq)

	if err != nil {
		return nil, errs.NewUnexceptedError(err.Message)
	}

	return &dto.NewReaderResponse{NewReaderId: int(dbResp.Id)}, nil

}

func NewReaderService(dbClient *pgx.Conn) ReaderService {
	return ReaderService{model.NewReaderRepositoryDb(dbClient)}
}

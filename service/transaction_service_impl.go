package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"tes-synapsis/helper"
	"tes-synapsis/model/api/transaction"
	"tes-synapsis/model/domain"
	"tes-synapsis/repository"
	"time"
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewTransactionService(transactionRepository repository.TransactionRepository, DB *sql.DB, validate *validator.Validate) TransactionService {
	return &TransactionServiceImpl{TransactionRepository: transactionRepository, DB: DB, Validate: validate}
}

func (service *TransactionServiceImpl) Create(ctx context.Context, request transaction.TransactionCreateRequest) transaction.TransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction := domain.Transaction{
		UserId:    request.UserId,
		Amount:    request.Amount,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	transaction = service.TransactionRepository.Save(ctx, tx, transaction)

	return helper.ToTransactionResponse(transaction)

}

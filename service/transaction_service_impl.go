package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator"
	"tes-synapsis/exception"
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
	UserRepository        repository.UserRepository
	PaymentService        PaymentService
}

func NewTransactionService(transactionRepository repository.TransactionRepository, DB *sql.DB, validate *validator.Validate, userRepository repository.UserRepository, paymentService PaymentService) TransactionService {
	return &TransactionServiceImpl{TransactionRepository: transactionRepository, DB: DB, Validate: validate, UserRepository: userRepository, PaymentService: paymentService}
}

func (service *TransactionServiceImpl) Create(ctx context.Context, request transaction.TransactionCreateRequest) transaction.TransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.Find(ctx, tx, ctx.Value("user").(string))
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if request.UserId != user.Id {
		panic(errors.New("user is not found"))
	}

	transaction := domain.Transaction{
		UserId:    request.UserId,
		Amount:    request.Amount,
		Status:    "pending",
		User:      user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	transaction = service.TransactionRepository.Save(ctx, tx, transaction)

	paymentUrl, err := service.PaymentService.GetPaymentUrl(ctx, transaction)
	helper.PanicIfError(err)

	transaction.PaymentURL = paymentUrl
	transaction = service.TransactionRepository.Update(ctx, tx, transaction)

	return helper.ToTransactionResponse(transaction)
}

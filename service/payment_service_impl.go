package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/veritrans/go-midtrans"
	"os"
	"strconv"
	"strings"
	"tes-synapsis/exception"
	"tes-synapsis/helper"
	"tes-synapsis/model/api/payment"
	"tes-synapsis/model/domain"
	"tes-synapsis/repository"
	"time"
)

type PaymentServiceImpl struct {
	DB                    *sql.DB
	Validate              *validator.Validate
	TransactionRepository repository.TransactionRepository
}

func NewPaymentService(DB *sql.DB, validate *validator.Validate, transactionRepository repository.TransactionRepository) PaymentService {
	return &PaymentServiceImpl{DB: DB, Validate: validate, TransactionRepository: transactionRepository}
}

func (service *PaymentServiceImpl) GetPaymentUrl(ctx context.Context, request domain.Transaction) (string, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("MIDCLIENT_SERVERKEY")
	midclient.ClientKey = os.Getenv("MIDCLIENT_CLIENTKEY")
	midclient.APIEnvType = midtrans.Sandbox

	var snapGateway midtrans.SnapGateway
	snapGateway = midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "order-" + strconv.Itoa(request.Id) + "-go-" + time.Stamp,
			GrossAmt: int64(request.Amount),
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: request.User.Username,
			Email: request.User.Email,
		},
	}

	token, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return token.RedirectURL, nil
}

func (service *PaymentServiceImpl) ProcessPayment(ctx context.Context, request payment.PaymentNotifikastionRequest) domain.Transaction {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderId := strings.Split(request.OrderId, "-")
	id, _ := strconv.Atoi(orderId[1])
	transactionId := id

	transaction, err := service.TransactionRepository.FindById(ctx, tx, transactionId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if request.PaymentType == "credit_card" && request.TransactionStatus == "capture" && request.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if request.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if request.TransactionStatus == "deny" || request.TransactionStatus == "expire" || request.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	transaction.UpdatedAt = time.Now()

	transaction = service.TransactionRepository.Update(ctx, tx, transaction)

	return transaction
}

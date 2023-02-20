package service

import (
	"context"
	"tes-synapsis/model/api/payment"
	"tes-synapsis/model/domain"
)

type PaymentService interface {
	GetPaymentUrl(ctx context.Context, request domain.Transaction) (string, error)
	ProcessPayment(ctx context.Context, request payment.PaymentNotifikastionRequest) domain.Transaction
}

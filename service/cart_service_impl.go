package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"tes-synapsis/helper"
	"tes-synapsis/model/api/cart"
	"tes-synapsis/model/domain"
	"tes-synapsis/repository"
)

type CartServiceImpl struct {
	CartRepository repository.CartRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewCartService(cartRepository repository.CartRepository, DB *sql.DB, validate *validator.Validate) CartService {
	return &CartServiceImpl{CartRepository: cartRepository, DB: DB, Validate: validate}
}

func (service *CartServiceImpl) Create(ctx context.Context, request cart.CartCreateRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cart := domain.Cart{
		ProductId: request.ProductId,
		Amount:    1,
	}

	service.CartRepository.Save(ctx, tx, cart)
}

func (service *CartServiceImpl) FindAll(ctx context.Context) []cart.CartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	carts := service.CartRepository.FindAll(ctx, tx)

	return helper.ToCartResponses(carts)
}

func (service *CartServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.CartRepository.Delete(ctx, tx, productId)
}

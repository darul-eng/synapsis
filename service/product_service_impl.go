package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"tes-synapsis/exception"
	"tes-synapsis/helper"
	"tes-synapsis/model/api/product"
	"tes-synapsis/model/domain"
	"tes-synapsis/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validete          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validete *validator.Validate) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository, DB: DB, Validete: validete}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request product.ProductCreateRequest) product.ProductResponse {
	err := service.Validete.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		Name:       request.Name,
		CategoryId: request.CategoryId,
		Price:      request.Price,
	}

	product = service.ProductRepository.Save(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindByCategoryId(ctx context.Context, categoryId int) []product.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products, err := service.ProductRepository.FindByCategoryId(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponses(products)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int) product.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(product)
}

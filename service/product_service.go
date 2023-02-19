package service

import (
	"context"
	"tes-synapsis/model/api/product"
)

type ProductService interface {
	Create(ctx context.Context, request product.ProductCreateRequest) product.ProductResponse
	FindByCategoryId(ctx context.Context, categoryId int) []product.ProductResponse
	FindById(ctx context.Context, productId int) product.ProductResponse
}

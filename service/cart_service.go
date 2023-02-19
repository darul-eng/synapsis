package service

import (
	"context"
	"tes-synapsis/model/api/cart"
)

type CartService interface {
	Create(ctx context.Context, request cart.CartCreateRequest)
	FindAll(ctx context.Context) []cart.CartResponse
	Delete(ctx context.Context, productId int)
}

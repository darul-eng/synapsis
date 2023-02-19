package helper

import (
	"tes-synapsis/model/api/auth"
	"tes-synapsis/model/api/cart"
	category "tes-synapsis/model/api/category"
	"tes-synapsis/model/api/product"
	"tes-synapsis/model/api/transaction"
	"tes-synapsis/model/domain"
)

func ToAuthResponse(data domain.User, token string) auth.AuthResponse {
	return auth.AuthResponse{
		Id:       data.Id,
		Username: data.Username,
		Token:    token,
	}
}

func ToCategoryResponse(data domain.Category) category.CategoryResponse {
	return category.CategoryResponse{
		Id:   data.Id,
		Name: data.Name,
	}
}

func ToProductResponse(data domain.Product) product.ProductResponse {
	return product.ProductResponse{
		Id:         data.Id,
		Name:       data.Name,
		Price:      data.Price,
		CategoryId: data.CategoryId,
	}
}

func ToProductResponses(datas []domain.Product) []product.ProductResponse {
	var products []product.ProductResponse
	for _, product := range datas {
		products = append(products, ToProductResponse(product))
	}

	return products
}

func ToCartResponse(data domain.Cart) cart.CartResponse {
	return cart.CartResponse{
		Id:        data.Id,
		ProductId: data.ProductId,
		Name:      data.Name,
		Amount:    data.Amount,
		Price:     data.Price,
		Category:  data.Category,
	}
}

func ToCartResponses(datas []domain.Cart) []cart.CartResponse {
	var carts []cart.CartResponse
	for _, cart := range datas {
		carts = append(carts, ToCartResponse(cart))
	}

	return carts
}

func ToTransactionResponse(data domain.Transaction) transaction.TransactionResponse {
	return transaction.TransactionResponse{
		Id:        data.Id,
		Amount:    data.Amount,
		Status:    data.Status,
		CreatedAt: data.CreatedAt,
	}
}

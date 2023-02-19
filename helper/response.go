package helper

import (
	"tes-synapsis/model/api/auth"
	category "tes-synapsis/model/api/category"
	"tes-synapsis/model/api/product"
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

package service

import (
	"context"
	user "tes-synapsis/model/api/category"
)

type CategoryService interface {
	Create(ctx context.Context, request user.CategoryCreateRequest) user.CategoryResponse
	FindById(ctx context.Context, categoryId int) user.CategoryResponse
}

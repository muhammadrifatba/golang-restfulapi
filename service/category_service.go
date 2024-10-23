package service

import (
	"context"
	"golang-restfulapi/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryI int)web.CategoryResponse
	FindAll(ctx context.Context)[]web.CategoryResponse
}
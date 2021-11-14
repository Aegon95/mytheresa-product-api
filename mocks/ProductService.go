// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/Aegon95/mytheresa-product-api/internal/db/entity"
	mock "github.com/stretchr/testify/mock"

	model "github.com/Aegon95/mytheresa-product-api/internal/model"
)

// ProductService is an autogenerated mock type for the ProductService type
type ProductService struct {
	mock.Mock
}

// GetProducts provides a mock function with given fields: ctx, category, than
func (_m *ProductService) GetProducts(ctx context.Context, category entity.Category, than int64) ([]model.ProductDTO, error) {
	ret := _m.Called(ctx, category, than)

	var r0 []model.ProductDTO
	if rf, ok := ret.Get(0).(func(context.Context, entity.Category, int64) []model.ProductDTO); ok {
		r0 = rf(ctx, category, than)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ProductDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.Category, int64) error); ok {
		r1 = rf(ctx, category, than)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

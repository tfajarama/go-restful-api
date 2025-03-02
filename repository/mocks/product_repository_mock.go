// Code generated by MockGen. DO NOT EDIT.
// Source: repository/product_repository.go
//
// Generated by this command:
//
//	mockgen -source=repository/product_repository.go -destination=repository/mocks/product_repository_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	"github.com/golang/mock/gomock"
	reflect "reflect"

	domain "github.com/aronipurwanto/go-restful-api/model/domain"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
	isgomock struct{}
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockProductRepository) Delete(ctx context.Context, product domain.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductRepositoryMockRecorder) Delete(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductRepository)(nil).Delete), ctx, product)
}

// FindAll mocks base method.
func (m *MockProductRepository) FindAll(ctx context.Context) ([]domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockProductRepositoryMockRecorder) FindAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockProductRepository)(nil).FindAll), ctx)
}

// FindById mocks base method.
func (m *MockProductRepository) FindById(ctx context.Context, productId uint64) (domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, productId)
	ret0, _ := ret[0].(domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockProductRepositoryMockRecorder) FindById(ctx, productId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockProductRepository)(nil).FindById), ctx, productId)
}

// Save mocks base method.
func (m *MockProductRepository) Save(ctx context.Context, product domain.Product) (domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, product)
	ret0, _ := ret[0].(domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockProductRepositoryMockRecorder) Save(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockProductRepository)(nil).Save), ctx, product)
}

// Update mocks base method.
func (m *MockProductRepository) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, product)
	ret0, _ := ret[0].(domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductRepositoryMockRecorder) Update(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductRepository)(nil).Update), ctx, product)
}

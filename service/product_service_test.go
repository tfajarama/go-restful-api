package service

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	mockValidator := validator.New()
	productService := NewProductService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.ProductCreateRequest
		mock      func()
		expect    web.ProductResponse
		expectErr bool
	}{
		{
			name: "success",
			input: web.ProductCreateRequest{
				Name:        "Wireless Mouse",
				Description: "A high-precision wireless mouse with ergonomic design.",
				Price:       29.99,
				StockQty:    150,
				SKU:         "WM12345",
				TaxRate:     0.07,
				CategoryID:  1,
				Category:    web.CategoryCreateRequest{Name: "Electronics"},
			},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Product{
					ProductID:   1,
					Name:        "Wireless Mouse",
					Description: "A high-precision wireless mouse with ergonomic design.",
					Price:       29.99,
					StockQty:    150,
					CategoryId:  1,
					SKU:         "WM12345",
					TaxRate:     0.07,
					Category:    domain.Category{Id: 1, Name: "Electronics"},
				}, nil)
			},
			expect: web.ProductResponse{
				Id:          1,
				Name:        "Wireless Mouse",
				Description: "A high-precision wireless mouse with ergonomic design.",
				Price:       29.99,
				StockQty:    150,
				SKU:         "WM12345",
				TaxRate:     0.07,
				CategoryID:  1,
				Category:    web.CategoryResponse{Id: 1, Name: "Electronics"},
			},
			expectErr: false,
		},
		{
			name:      "validation error",
			input:     web.ProductCreateRequest{Name: ""},
			mock:      func() {},
			expect:    web.ProductResponse{},
			expectErr: true,
		},
		{
			name: "repository error",
			input: web.ProductCreateRequest{
				Name:        "Wireless Mouse",
				Description: "A high-precision wireless mouse with ergonomic design.",
				Price:       29.99,
				StockQty:    150,
				SKU:         "WM12345",
				TaxRate:     0.07,
				CategoryID:  1,
				Category:    web.CategoryCreateRequest{Name: "Electronics"},
			},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Product{}, errors.New("database error"))
			},
			expect:    web.ProductResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := productService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := NewProductService(mockRepo, validator.New())

	tests := []struct {
		name      string
		productId uint64
		mock      func()
		expectErr bool
	}{
		{
			name:      "success",
			productId: 1,
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), uint64(1)).Return(domain.Product{
					ProductID:   1,
					Name:        "Wireless Mouse",
					Description: "A high-precision wireless mouse with ergonomic design.",
					Price:       29.99,
					StockQty:    150,
					SKU:         "WM12345",
					TaxRate:     0.07,
					CategoryId:  1,
					Category:    domain.Category{Id: 1, Name: "Electronics"},
				}, nil)
				mockRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectErr: false,
		},
		{
			name:      "not found",
			productId: 99,
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), uint64(99)).Return(domain.Product{}, errors.New("not found"))
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := productService.Delete(context.Background(), tt.productId)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mockProductRepo *mocks.MockProductRepository)
		input   web.ProductUpdateRequest
		expects error
	}{
		{
			name: "Success",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				mockProductRepo.EXPECT().FindById(gomock.Any(), uint64(1)).
					Return(domain.Product{
						ProductID:   1,
						Name:        "Wireless Mouse",
						Description: "A high-precision wireless mouse with ergonomic design.",
						Price:       29.99,
						StockQty:    150,
						SKU:         "WM12345",
						TaxRate:     0.07,
						CategoryId:  1,
						Category:    domain.Category{Id: 1, Name: "Electronics"},
					}, nil)
				mockProductRepo.EXPECT().Update(gomock.Any(), gomock.Any()).
					Return(domain.Product{
						ProductID:   1,
						Name:        "[NEW] Wireless Mouse",
						Description: "A high-precision wireless mouse with ergonomic design.",
						Price:       29.99,
						StockQty:    150,
						SKU:         "WM12345",
						TaxRate:     0.07,
						CategoryId:  1,
						Category:    domain.Category{Id: 1, Name: "Electronics"},
					}, nil)
			},
			input: web.ProductUpdateRequest{
				Id:          1,
				Name:        "[NEW] Wireless Mouse",
				Description: "A high-precision wireless mouse with ergonomic design.",
				Price:       29.99,
				StockQty:    150,
				SKU:         "WM12345",
				TaxRate:     0.07,
				CategoryID:  1,
				Category:    web.CategoryUpdateRequest{Id: 1, Name: "Electronics"},
			},
			expects: nil,
		},
		{
			name: "Not Found",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				mockProductRepo.EXPECT().FindById(gomock.Any(), uint64(1)).
					Return(domain.Product{}, errors.New("not found"))
			},
			input: web.ProductUpdateRequest{
				Id:          1,
				Name:        "[NEW] Wireless Mouse",
				Description: "A high-precision wireless mouse with ergonomic design.",
				Price:       29.99,
				StockQty:    150,
				SKU:         "WM12345",
				TaxRate:     0.07,
				CategoryID:  1,
				Category:    web.CategoryUpdateRequest{Id: 1, Name: "Electronics"},
			},
			expects: errors.New("not found"),
		},
		{
			name: "Validation Error - Empty Name",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				// Tidak perlu mock FindById karena validasi gagal sebelum ke repository
			},
			input: web.ProductUpdateRequest{
				Id:          1,
				Name:        "",
				Description: "A high-precision wireless mouse with ergonomic design.",
				Price:       29.99,
				StockQty:    150,
				SKU:         "WM12345",
				TaxRate:     0.07,
				CategoryID:  1,
				Category:    web.CategoryUpdateRequest{Id: 1, Name: "Electronics"},
			},
			expects: errors.New("ProductUpdateRequest.Name"),
		},
		{
			name: "Database Error on Update",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				mockProductRepo.EXPECT().FindById(gomock.Any(), uint64(1)).
					Return(domain.Product{
						ProductID:   1,
						Name:        "Wireless Mouse",
						Description: "A high-precision wireless mouse with ergonomic design.",
						Price:       29.99,
						StockQty:    150,
						SKU:         "WM12345",
						TaxRate:     0.07,
						CategoryId:  1,
						Category:    domain.Category{Id: 1, Name: "Electronics"},
					}, nil)
				mockProductRepo.EXPECT().Update(gomock.Any(), gomock.Any()).
					Return(domain.Product{}, errors.New("database error"))
			},
			input: web.ProductUpdateRequest{
				Id:          1,
				Name:        "[NEW] Wireless Mouse",
				Description: "A high-precision wireless mouse with ergonomic design.",
				Price:       29.99,
				StockQty:    150,
				SKU:         "WM12345",
				TaxRate:     0.07,
				CategoryID:  1,
				Category:    web.CategoryUpdateRequest{Id: 1, Name: "Electronics"},
			},
			expects: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProductRepo := mocks.NewMockProductRepository(ctrl)
			tt.mock(mockProductRepo)

			service := NewProductService(mockProductRepo, validator.New())
			_, err := service.Update(context.Background(), tt.input)

			if tt.expects != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expects.Error()) // Alternatif untuk assert.ErrorContains
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFindAllProducts(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mockProductRepo *mocks.MockProductRepository)
		expects []web.ProductResponse
		err     error
	}{
		{
			name: "Success",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				mockProductRepo.EXPECT().FindAll(gomock.Any()).Return([]domain.Product{
					{
						ProductID:   1,
						Name:        "Wireless Mouse",
						Description: "A high-precision wireless mouse with ergonomic design.",
						Price:       29.99,
						StockQty:    150,
						SKU:         "WM12345",
						TaxRate:     0.07,
						CategoryId:  1,
						Category:    domain.Category{Id: 1, Name: "Electronics"},
					},
					{
						ProductID:   2,
						Name:        "Bluetooth Headphones",
						Description: "Noise-cancelling over-ear headphones with long battery life.",
						Price:       89.99,
						StockQty:    75,
						SKU:         "BH67890",
						TaxRate:     0.07,
						CategoryId:  1,
						Category:    domain.Category{Id: 1, Name: "Electronics"},
					},
					{
						ProductID:   3,
						Name:        "Smartphone Stand",
						Description: "Adjustable stand for smartphones and tablets.",
						Price:       15.99,
						StockQty:    200,
						SKU:         "SS11223",
						TaxRate:     0.05,
						CategoryId:  2,
						Category:    domain.Category{Id: 2, Name: "Accessories"},
					},
					{
						ProductID:   4,
						Name:        "USB-C Charger",
						Description: "Fast-charging USB-C charger with multiple ports.",
						Price:       24.99,
						StockQty:    120,
						SKU:         "UC33445",
						TaxRate:     0.07,
						CategoryId:  1,
						Category:    domain.Category{Id: 1, Name: "Electronics"},
					},
					{
						ProductID:   5,
						Name:        "Gaming Keyboard",
						Description: "Mechanical keyboard with customizable RGB lighting.",
						Price:       59.99,
						StockQty:    80,
						SKU:         "GK55667",
						TaxRate:     0.07,
						CategoryId:  1,
						Category:    domain.Category{Id: 1, Name: "Electronics"},
					}}, nil)
			},
			expects: []web.ProductResponse{{
				Id:          1,
				Name:        "Wireless Mouse",
				Description: "A high-precision wireless mouse with ergonomic design.",
				Price:       29.99,
				StockQty:    150,
				SKU:         "WM12345",
				TaxRate:     0.07,
				CategoryID:  1,
				Category:    web.CategoryResponse{Id: 1, Name: "Electronics"},
			},
				{
					Id:          2,
					Name:        "Bluetooth Headphones",
					Description: "Noise-cancelling over-ear headphones with long battery life.",
					Price:       89.99,
					StockQty:    75,
					SKU:         "BH67890",
					TaxRate:     0.07,
					CategoryID:  1,
					Category:    web.CategoryResponse{Id: 1, Name: "Electronics"},
				},
				{
					Id:          3,
					Name:        "Smartphone Stand",
					Description: "Adjustable stand for smartphones and tablets.",
					Price:       15.99,
					StockQty:    200,
					SKU:         "SS11223",
					TaxRate:     0.05,
					CategoryID:  2,
					Category:    web.CategoryResponse{Id: 2, Name: "Accessories"},
				},
				{
					Id:          4,
					Name:        "USB-C Charger",
					Description: "Fast-charging USB-C charger with multiple ports.",
					Price:       24.99,
					StockQty:    120,
					SKU:         "UC33445",
					TaxRate:     0.07,
					CategoryID:  1,
					Category:    web.CategoryResponse{Id: 1, Name: "Electronics"},
				},
				{
					Id:          5,
					Name:        "Gaming Keyboard",
					Description: "Mechanical keyboard with customizable RGB lighting.",
					Price:       59.99,
					StockQty:    80,
					SKU:         "GK55667",
					TaxRate:     0.07,
					CategoryID:  1,
					Category:    web.CategoryResponse{Id: 1, Name: "Electronics"},
				}},
			err: nil,
		},
		{
			name: "Database Error",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				mockProductRepo.EXPECT().FindAll(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expects: nil,
			err:     errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProductRepo := mocks.NewMockProductRepository(ctrl)
			tt.mock(mockProductRepo)

			service := NewProductService(mockProductRepo, validator.New())
			result, err := service.FindAll(context.Background())
			assert.Equal(t, tt.expects, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestFindByIdProduct(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mockProductRepo *mocks.MockProductRepository)
		input   uint64
		expects web.ProductResponse
		err     error
	}{
		{
			name: "Success",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				mockProductRepo.EXPECT().FindById(gomock.Any(), uint64(1)).Return(domain.Product{
					ProductID:   1,
					Name:        "Wireless Mouse",
					Description: "A high-precision wireless mouse with ergonomic design.",
					Price:       29.99,
					StockQty:    150,
					SKU:         "WM12345",
					TaxRate:     0.07,
					CategoryId:  1,
					Category:    domain.Category{Id: 1, Name: "Electronics"},
				}, nil)
			},
			input: 1,
			expects: web.ProductResponse{
				Id:          1,
				Name:        "Wireless Mouse",
				Description: "A high-precision wireless mouse with ergonomic design.",
				Price:       29.99,
				StockQty:    150,
				SKU:         "WM12345",
				TaxRate:     0.07,
				CategoryID:  1,
				Category:    web.CategoryResponse{Id: 1, Name: "Electronics"},
			},
			err: nil,
		},
		{
			name: "Not Found",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				mockProductRepo.EXPECT().FindById(gomock.Any(), uint64(1)).Return(domain.Product{}, errors.New("not found"))
			},
			input:   1,
			expects: web.ProductResponse{},
			err:     errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProductRepo := mocks.NewMockProductRepository(ctrl)
			tt.mock(mockProductRepo)

			service := NewProductService(mockProductRepo, validator.New())
			result, err := service.FindById(context.Background(), tt.input)
			assert.Equal(t, tt.expects, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

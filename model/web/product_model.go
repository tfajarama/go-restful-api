package web

type ProductCreateRequest struct {
	Name        string                `json:"name" validate:"required,max=32,min=10"`
	Description string                `json:"description"`
	Price       float64               `json:"price" validate:"required,gte=0"`
	StockQty    int                   `json:"stock_qty" validate:"required,gte=0"`
	SKU         string                `json:"sku" validate:"required"`
	TaxRate     float64               `json:"tax_rate" validate:"required,gte=0"`
	CategoryID  uint64                `json:"category_id" validate:"required"`
	Category    CategoryCreateRequest `json:"category" validate:"required"`
}

type ProductUpdateRequest struct {
	Id          uint64                `json:"id" validate:"required,gte=0"`
	Name        string                `json:"name" validate:"required,max=32,min=10"`
	Description string                `json:"description"`
	Price       float64               `json:"price" validate:"required,gte=0"`
	StockQty    int                   `json:"stock_qty" validate:"required,gte=0"`
	SKU         string                `json:"sku" validate:"required"`
	TaxRate     float64               `json:"tax_rate" validate:"required,gte=0"`
	CategoryID  uint64                `json:"category_id" validate:"required"`
	Category    CategoryUpdateRequest `json:"category" validate:"required"`
}

type ProductResponse struct {
	Id          uint64           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       float64          `json:"price"`
	StockQty    int              `json:"stock_qty"`
	SKU         string           `json:"sku"`
	TaxRate     float64          `json:"tax_rate"`
	CategoryID  uint64           `json:"category_id"`
	Category    CategoryResponse `json:"category"`
}

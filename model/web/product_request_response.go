package web

type ProductCreateRequest struct {
	Name        string  `json:"product_name" validate:"required,min=1,max=100"`
	Description string  `json:"product_description" validate:"required,min=1,max=255"`
	Price       float64 `json:"product_price" validate:"required,min=1"`
	StockQty    int64   `json:"product_stock_qty" validate:"required,min=0"`
	CategoryID  int64   `json:"category_id" validate:"required"`
	SKU         string  `json:"product_sku" validate:"required,min=1,max=100"`
	TaxRate     float64 `json:"product_tax_rate" validate:"required,min=0"`
}

type ProductUpdateRequest struct {
	ProductID   string  `json:"product_id" validate:"required"`
	Name        string  `json:"product_name" validate:"required,min=1,max=100"`
	Description string  `json:"product_description" validate:"required,min=1,max=255"`
	Price       float64 `json:"product_price" validate:"required,min=1"`
	StockQty    int64   `json:"product_stock_qty" validate:"required,min=0"`
	CategoryID  int64   `json:"category_id" validate:"required"`
	SKU         string  `json:"product_sku" validate:"required,min=1,max=100"`
	TaxRate     float64 `json:"product_tax_rate" validate:"required,min=0"`
}

type ProductResponse struct {
	ProductID   string  `json:"product_id"`
	Name        string  `json:"product_name"`
	Description string  `json:"product_description"`
	Price       float64 `json:"product_price"`
	StockQty    int64   `json:"product_stock_qty"`
	CategoryID  int64   `json:"category_id"`
	SKU         string  `json:"product_sku"`
	TaxRate     float64 `json:"product_tax_rate"`
}

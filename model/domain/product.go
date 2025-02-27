package domain

type Product struct {
	ProductID   string  `json:"product_id" gorm:"primaryKey; unique; column:product_id"`
	Name        string  `json:"product_name" gorm:"column:product_name; type:varchar(100)"`
	Description string  `json:"product_description" gorm:"column:product_description; type:varchar(255)"`
	Price       float64 `json:"product_price" gorm:"column:product_price; type:float"`
	StockQty    int64   `json:"product_stock_qty" gorm:"column:stock_qty; type:int"`
	CategoryID  int64   `json:"category_id" gorm:"column:category_id; type:int"`
	SKU         string  `json:"product_sku" gorm:"column:product_sku; type:varchar(100)"`
	TaxRate     float64 `json:"product_tax_rate" gorm:"column:tax_rate; type:float"`
}

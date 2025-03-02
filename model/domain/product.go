package domain

type Product struct {
	ProductID   uint64   `gorm:"primary_key;autoIncrement;column:id"`
	Name        string   `gorm:"column:product_name; length:255"`
	Description string   `gorm:"column:product_description; length:255"`
	Price       float64  `gorm:"column:product_price"`
	StockQty    int      `gorm:"column:stock_qty"`
	SKU         string   `gorm:"column:product_sku"`
	TaxRate     float64  `gorm:"column:tax_rate"`
	CategoryId  uint64   `gorm:"column:category_id"`
	Category    Category `gorm:"foreign_key:CategoryId;references:Id"`
}

type ProductError struct {
	Product Product `json:"product"`
	Error   error   `json:"error"`
}

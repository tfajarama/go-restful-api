package domain

type Category struct {
	Id      uint64    `gorm:"primary_key;autoIncrement;column:id"`
	Name    string    `gorm:"column:name"`
	Product []Product `gorm:"foreign_key:CategoryId;references:Id"`
}

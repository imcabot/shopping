package product

import (
	"gorm.io/gorm"
	"shopping/category"
)

//商品结构体
type Product struct {
	gorm.Model
	Name       string
	SKU        string
	Desc       string
	StockCount int
	Price      float32
	CategoryID uint
	Category   category.Category `json:"-"`
	IsDeleted  bool
}

func NewProduct(name string, desc string, stockCount int, price float32, cid uint) *Product {
	return &Product{
		Name:       name,
		Desc:       desc,
		StockCount: stockCount,
		Price:      price,
		CategoryID: cid,
		IsDeleted:  false,
	}
}

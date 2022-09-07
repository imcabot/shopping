package order

import (
	"gorm.io/gorm"
	"shopping/domain/product"
	"shopping/domain/user"
)

type Order struct {
	gorm.Model
	UserID       uint
	User         user.User     `gorm:"foreignKey:ID;references:UserID" json:"-"`
	OrderedItems []OrderedItem `gorm:"foreignKey:OrderID"`
	TotalPrice   float32
	IsCanceled   bool //是否退单
}

type OrderedItem struct {
	gorm.Model
	Product    product.Product `gorm:"foreignKey:ProductID"`
	ProductID  uint
	Count      int
	OrderID    uint
	IsCanceled bool //是否退单
}

func NewOrder(uid uint, items []OrderedItem) *Order {
	var totalPrice float32
	for _, item := range items {
		totalPrice += item.Product.Price
	}
	return &Order{
		UserID:       uid,
		OrderedItems: items,
		TotalPrice:   totalPrice,
		IsCanceled:   false,
	}
}

func NewOrderItem(count int, pid uint) *OrderedItem {
	return &OrderedItem{
		Count:      count,
		ProductID:  pid,
		IsCanceled: false,
	}
}

package cart

import (
	"gorm.io/gorm"
	"shopping/domain/product"
	"shopping/domain/user"
)

type Cart struct {
	gorm.Model
	UserID uint
	User   user.User `gorm:"foreignKey:ID;references:UserID"`
}

func NewCart(uid uint) *Cart {
	return &Cart{
		UserID: uid,
	}
}

type Item struct {
	gorm.Model
	Product   product.Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Count     int
	CartID    uint
	Cart      Cart `gorm:"foreignKey:CartID" json:"-"`
}

func NewCartItem(productId uint, cartID uint, count int) *Item {
	return &Item{
		ProductID: productId,
		CartID:    cartID,
		Count:     count,
	}
}

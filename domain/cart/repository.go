package cart

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

//创建表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Cart{})
	if err != nil {
		log.Println(err)
	}
}

//更新
func (r *Repository) Update(cart Cart) error {
	result := r.db.Save(cart)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//根据用户ID查找或创建购物车
func (r *Repository) FindOrCreateByUserId(userID uint) (*Cart, error) {
	var cart *Cart
	err := r.db.Where(Cart{UserID: userID}).Attrs(NewCart(userID)).FirstOrCreate(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

//根据用户ID查找购物车
func (r *Repository) FindbyUserID(userID uint) (*Cart, error) {
	var cart *Cart
	err := r.db.Where(Cart{UserID: userID}).Attrs(NewCart(userID)).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

type ItemRepository struct {
	db *gorm.DB
}

//实例化Item表
func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

// 生成Item表
func (r *ItemRepository) Migration() {
	err := r.db.AutoMigrate(&Item{})
	if err != nil {
		log.Println(err)
	}
}

//更新Item
func (r *ItemRepository) Update(item Item) error {
	result := r.db.Save(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//根据商品ID和购物车ID查找Item
func (r *ItemRepository) FindByID(pid uint, cid uint) (*Item, error) {
	var item *Item
	err := r.db.Where(&Item{ProductID: pid, CartID: cid}).First(&item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

//创建Item
func (r *ItemRepository) Create(item *Item) error {
	result := r.db.Create(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//返回购物车中所有item
func (r *ItemRepository) GetItems(cartId uint) ([]Item, error) {
	var cartItems []Item
	err := r.db.Where(Item{CartID: cartId}).Find(&cartItems).Error
	if err != nil {
		return nil, err
	}
	for i, item := range cartItems {
		err := r.db.Model(item).Association("Product").Find(&cartItems[i].Product)
		if err != nil {
			return nil, err
		}
	}
	return cartItems, nil

}

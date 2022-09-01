package order

import (
	"gorm.io/gorm"
	"log"
)

type OrderedItemRepository struct {
	db *gorm.DB
}

func NewOrderedItemRepository(db *gorm.DB) *OrderedItemRepository {
	return &OrderedItemRepository{
		db: db,
	}
}

func (r *OrderedItemRepository) Migration() {
	err := r.db.AutoMigrate(&OrderedItem{})
	if err != nil {
		log.Println(err)
	}
}

func (r *OrderedItemRepository) Update(item OrderedItem) error {
	result := r.db.Save(&item)
	if result.Error != nil {
		return result.Error

	}
	return nil
}

func (r *OrderedItemRepository) Create(ci *OrderedItem) error {
	result := r.db.Create(ci)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

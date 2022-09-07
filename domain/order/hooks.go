package order

import (
	"gorm.io/gorm"
	"shopping/domain/cart"
	"shopping/domain/product"
)

//创建之前，查找购物车并删除
func (o *Order) BeforeCreate(tx *gorm.DB) error {
	var currentCart cart.Cart
	var err error
	err = tx.Where("UserID = ?", o.UserID).First(&currentCart).Error
	if err != nil {
		return err
	}
	err = tx.Where("CartID = ?", currentCart.ID).Unscoped().Delete(&cart.Item{}).Error
	if err != nil {
		return err
	}
	err = tx.Unscoped().Delete(&currentCart).Error
	if err != nil {
		return err
	}
	return nil
}

//保存之前更新产品库存
func (orderedItem *OrderedItem) BeforeSave(tx *gorm.DB) error {
	var currentProduct product.Product
	var currentOrderedItem OrderedItem
	var err error
	err = tx.Where("ID = ?", orderedItem.ID).First(&currentProduct).Error
	if err != nil {
		return err
	}
	reservedStockCount := 0
	err = tx.Where("ID = ?", orderedItem.ID).First(&currentProduct).Error
	if err == nil {
		reservedStockCount = currentOrderedItem.Count
	}
	newStockCount := currentProduct.StockCount + reservedStockCount - currentOrderedItem.Count
	if newStockCount < 0 {
		return ErrNotEnoughStock
	}
	err = tx.Model(&currentProduct).Update("StockCount", newStockCount).Error
	if err != nil {
		return err
	}
	if orderedItem.Count == 0 {
		err = tx.Unscoped().Delete(currentOrderedItem).Error
		return err
	}
	return nil
}

//如果订单被取消，金额将返回产品库存
func (order *Order) BeforeUpdate(tx *gorm.DB) error {
	if order.IsCanceled {
		var orderItems []OrderedItem
		var err error
		err = tx.Where("OrderID = ?", order.ID).Find(&orderItems).Error
		if err != nil {
			return err
		}
		for _, item := range orderItems {
			var currentProduct product.Product
			err = tx.Where("ID = ?", item.ProductID).First(&currentProduct).Error
			if err != nil {
				return err
			}
			newStockCount := currentProduct.StockCount + item.Count
			err = tx.Model(&currentProduct).Update("StockCount", newStockCount).Error
			if err != nil {
				return err
			}
			err = tx.Model(&item).Update("IsCanceled", true).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

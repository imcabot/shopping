package order

import (
	"shopping/domain/cart"
	"shopping/domain/product"
	"shopping/utils/pagination"
	"time"
)

var day14ToHours float64 = 336

type Service struct {
	orderRepository       Repository
	orderedItemRepository OrderedItemRepository
	productRepository     product.Repository
	cartRepository        cart.Repository
	cartItemRepository    cart.ItemRepository
}

func NewService(orderRepository Repository,
	orderedItemRepository OrderedItemRepository,
	productRepository product.Repository,
	cartRepository cart.Repository,
	cartItemRepository cart.ItemRepository) *Service {
	orderRepository.Migration()
	orderedItemRepository.Migration()
	return &Service{
		orderRepository:       orderRepository,
		orderedItemRepository: orderedItemRepository,
		productRepository:     productRepository,
		cartRepository:        cartRepository,
		cartItemRepository:    cartItemRepository,
	}

}

//完成订单
func (s *Service) CompleteOrder(userID uint) error {
	currentCart, err := s.cartRepository.FindOrCreateByUserId(userID)
	if err != nil {
		return err
	}
	cartItems, err := s.cartItemRepository.GetItems(currentCart.UserID)
	if err != nil {
		return err
	}
	if len(cartItems) == 0 {
		return ErrEmptyCartFond
	}
	orderedItems := make([]OrderedItem, 0)
	for _, item := range cartItems {
		orderedItems = append(orderedItems, *NewOrderItem(item.Count, item.ProductID))
	}
	err = s.orderRepository.Create(NewOrder(userID, orderedItems))
	if err != nil {
		return err
	}
	return nil
}

//取消订单
func (s *Service) CancelOrder(uid, oid uint) error {
	currentOrder, err := s.orderRepository.FindByOrderId(oid)
	if err != nil {
		return err
	}
	if currentOrder.UserID != uid {
		return ErrInvalidOrderID
	}
	if currentOrder.CreatedAt.Sub(time.Now()).Hours() > day14ToHours {
		return ErrCancelDurationPassed
	}
	currentOrder.IsCanceled = true
	err = s.orderRepository.Update(*currentOrder)
	if err != nil {
		return err
	}
	return nil
}

//获得订单
func (s *Service) GetAll(page *pagination.Pages, uid uint) *pagination.Pages {
	orders, count := s.orderRepository.GetAll(page.Page, page.PageSize, uid)
	page.Items = orders
	page.TotalCount = count
	return page
}

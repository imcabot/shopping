package product

import (
	"shopping/utils/pagination"
)

type Service struct {
	productRepository Repository
}

func NewService(productRepostory Repository) *Service {
	productRepostory.Migration()
	return &Service{
		productRepository: productRepostory,
	}
}

func (s *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	products, count := s.productRepository.GetAll(page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}

func (s *Service) CreateProduct(name string, desc string, count int, price float32, cid uint) error {
	newProduct := NewProduct(name, desc, count, price, cid)
	err := s.productRepository.Create(newProduct)
	return err
}

func (s *Service) DeleteProduct(sku string) error {
	err := s.productRepository.Delete(sku)
	return err
}

func (s *Service) UpdateProduct(product *Product) error {
	err := s.productRepository.Update(*product)
	return err
}

func (s *Service) SearchProduct(text string, page *pagination.Pages) *pagination.Pages {
	products, count := s.productRepository.SearchByString(text, page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}

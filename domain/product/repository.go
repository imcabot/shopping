package product

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

//type Dd interface {
//	Migration()
//	Update(updateProduct Product) error
//	SearchByString(str string, pageIndex, pageSize int) ([]Product, int)
//	FindBySku(sku string) (*Product, error)
//	Create(p *Product) error
//	GetAll(pageIndex, pageSize int) ([]Product, int)
//	Delete(sku string) error
//}

func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Product{})
	if err != nil {
		log.Print(err)
	}
}

func (r *Repository) Update(updateProduct Product) error {
	saveProduct, err := r.FindBySku(updateProduct.SKU)
	if err != nil {
		return err
	}
	err = r.db.Model(&saveProduct).Updates(updateProduct.SKU).Error
	return err
}

func (r *Repository) SearchByString(str string, pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	convertedStr := "%" + str + "%"
	var count int64
	r.db.Where("IsDeleted = ?", false).Where(
		"Name like ? or Sku like ?", convertedStr,
		convertedStr).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)
	return products, int(count)
}

func (r *Repository) FindBySku(sku string) (*Product, error) {
	var product *Product
	err := r.db.Where("IsDeleted = ?", 0).Where(Product{SKU: sku}).First(&product).Error
	if err != nil {
		return nil, ErrProductNotFound
	}
	return product, nil
}

func (r *Repository) Create(p *Product) error {
	result := r.db.Create(p)
	return result.Error
}

func (r *Repository) GetAll(pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	var count int64

	r.db.Where("IsDeleted = ?", 0).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)
	return products, int(count)
}

func (r *Repository) Delete(sku string) error {
	currentProduct, err := r.FindBySku(sku)
	if err != nil {
		return err
	}
	currentProduct.IsDeleted = true
	err = r.db.Save(currentProduct).Error
	return err
}

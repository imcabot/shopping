package category

import (
	"mime/multipart"
	"shoping/utils/csv_helper"
	"shoping/utils/pagination"
)

type Service struct {
	r Repository
}

//type Dd interface {
//	Create(category *Category) error
//	BulkCreate(fileHeader multipart.FileHeader) (int, error)
//	Get(page *pagination.Pages) *pagination.Pages
//}

//实例化商品分类service
func NewCategoryservice(r Repository) *Service {
	//生成表
	r.Migration()
	//插入测试数据
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

//创建分类
func (s *Service) Create(category *Category) error {
	existCity := s.r.GetByName(category.Name)
	if len(existCity) > 0 {
		return ErrCategoryExistWithName
	}
	err := s.r.Create(category)
	if err != nil {
		return err
	}
	return nil
}

//批量创建分类
func (s *Service) BulkCreate(fileHeader *multipart.FileHeader) (int, error) {
	categories := make([]*Category, 0)
	bulkCategory, err := csv_helper.ReadCsv(fileHeader)
	if err != nil {
		return 0, err
	}
	for _, categoryVariables := range bulkCategory {
		categories = append(categories, NewCategory(categoryVariables[0], categoryVariables[1]))
	}
	count, err := s.r.BulCreate(categories)
	if err != nil {
		return count, err
	}
	return count, nil
}

//获得分页商品分类
func (s *Service) Get(page *pagination.Pages) *pagination.Pages {
	categories, count := s.r.GetAll(page.Page, page.PageSize)
	page.Items = categories
	page.TotalCount = count
	return page
}

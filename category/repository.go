package category

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

//新建商品分类
func NewCategoryRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

//生成商品分类表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Category{})
	if err != nil {
		log.Print(err)
	}
}

//生成商品分类测试数据
func (r *Repository) InsertSampleData() {
	categorys := []Category{
		{Name: "CAT1", Desc: "Category1"},
		{Name: "CAT2", Desc: "Category2"},
	}
	for _, c := range categorys {
		r.db.Where(Category{Name: c.Name}).Attrs(Category{Name: c.Name}).FirstOrCreate(&c)
	}
}

//创建商品分类
func (r *Repository) Create(c *Category) error {
	result := r.db.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//通过名称查询商品分类
func (r *Repository) GetByName(name string) []Category {
	var categorys []Category
	r.db.Where("Name = ?", name).Find(&categorys)
	return categorys
}

//批量创建商品分类
func (r *Repository) BulCreate(categorys []*Category) (int, error) {
	var count int64
	err := r.db.Create(&categorys).Count(&count).Error
	return int(count), err
}

//获得分页商品分类
func (r *Repository) GetAll(pageIndex int, pageSize int) ([]Category, int) {
	var categorys []Category
	var count int64
	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categorys).Count(&count)
	return categorys, int(count)
}

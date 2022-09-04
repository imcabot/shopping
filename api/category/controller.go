package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shoping/category"
	"shoping/utils/api_helper"
	"shoping/utils/pagination"
)

type Controller struct {
	categoryService *category.Service
}

func NewCategoryController(categoryService *category.Service) *Controller {
	return &Controller{
		categoryService: categoryService,
	}
}

func (c *Controller) CreateCategory(g *gin.Context) {
	var req CreateCategoryRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return

	}
	newCategory := category.NewCategory(req.Name, req.Desc)
	err := c.categoryService.Create(newCategory)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(
		http.StatusCreated, api_helper.Response{
			Message: "Category created",
		})
}

func (c *Controller) BulkCreateCategory(g *gin.Context) {
	fileHeader, err := g.FormFile("file")
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	count, err := c.categoryService.BulkCreate(fileHeader)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{
		Message: fmt.Sprintf("'%s' uploaded! '%d' new categories created", fileHeader.Filename, count),
	})
}

func (c *Controller) GetCategories(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	page = c.categoryService.GetAll(page)
	g.JSON(http.StatusOK, page)
}

package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoping/domain/product"
	"shoping/utils/api_helper"
	"shoping/utils/pagination"
)

type Controller struct {
	productService *product.Service
}

func NewProductController(productService *product.Service) *Controller {
	return &Controller{productService: productService}
}

func (c *Controller) GetProducts(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	queryText := g.Query("qt")
	if queryText != "" {
		page = c.productService.SearchProduct(queryText, page)

	} else {
		page = c.productService.GetAll(page)
	}
	g.JSON(http.StatusOK, page)
}

func (c *Controller) CreateProduct(g *gin.Context) {
	var req CreateProductRequest
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	err = c.productService.CreateProduct(req.Name, req.Desc, req.Count, req.Price, req.CategoryID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{
		Message: "Product Created",
	})

}

func (c *Controller) DeleteProduct(g *gin.Context) {
	var req DeleteProductRequest
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	err = c.productService.DeleteProduct(req.SKU)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{
		Message: "Product Deleted",
	})

}

func (c *Controller) UpdateProduct(g *gin.Context) {
	var req UpdateProductRequest
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	err = c.productService.UpdateProduct(req.ToProduct())
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{
		Message: "Product Updated",
	})

}

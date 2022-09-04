package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoping/domain/cart"
	"shoping/utils/api_helper"
)

type Controller struct {
	cartService *cart.Service
}

func NewCartController(cartService *cart.Service) *Controller {
	return &Controller{cartService: cartService}
}

func (c *Controller) AddItem(g *gin.Context) {
	var req ItemCartRequest
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	userID := api_helper.GetUserId(g)
	err = c.cartService.AddItem(userID, req.SKU, req.Count)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{
		Message: "Item added to cart",
	})
}

func (c *Controller) UpdateItem(g *gin.Context) {
	var req ItemCartRequest
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	userID := api_helper.GetUserId(g)
	err = c.cartService.UpdateItem(userID, req.SKU, req.Count)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{
		Message: "updated",
	})

}

func (c *Controller) GetCart(g *gin.Context) {
	userID := api_helper.GetUserId(g)
	result, err := c.cartService.GetCartItem(userID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, result)
}

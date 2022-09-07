package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping/domain/cart"
	"shopping/utils/api_helper"
)

type Controller struct {
	cartService *cart.Service
}

func NewCartController(cartService *cart.Service) *Controller {
	return &Controller{cartService: cartService}
}

// AddItem godoc
// @Summary 添加Item
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param ItemCartRequest body ItemCartRequest true "product information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart/item [post]
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

// UpdateItem godoc
// @Summary 更新Item
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param ItemCartRequest body ItemCartRequest true "product information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart/item [patch]
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

// GetCart godoc
// @Summary 获得购物车商品列表
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Success 200 {array} cart.Item
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart [get]
func (c *Controller) GetCart(g *gin.Context) {
	userID := api_helper.GetUserId(g)
	result, err := c.cartService.GetCartItem(userID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, result)
}

package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoping/domain/order"
	"shoping/utils/api_helper"
	"shoping/utils/pagination"
)

type Controller struct {
	orderService *order.Service
}

func NewOrderController(orderService *order.Service) *Controller {
	return &Controller{orderService: orderService}
}

// CompleteOrder godoc
// @Summary 完成订单
// @Tags Order
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /order [post]
func (c *Controller) CompleteOrder(g *gin.Context) {
	userID := api_helper.GetUserId(g)
	err := c.orderService.CompleteOrder(userID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusCreated, api_helper.Response{Message: "Order Created"})
}

// CancelOrder godoc
// @Summary 取消订单
// @Tags Order
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param CancelOrderRequest body CancelOrderRequest true "order information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /order [delete]
func (c *Controller) CancelOrder(g *gin.Context) {
	var req CancelOrderRequest
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	userID := api_helper.GetUserId(g)
	err = c.orderService.CancelOrder(userID, req.OrderID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusCreated, api_helper.Response{Message: "Order Canceled"})
}

// GetOrders godoc
// @Summary 获得订单列表
// @Tags Order
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} pagination.Pages
// @Router /order [get]
func (c *Controller) GetOrders(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	userID := api_helper.GetUserId(g)
	page = c.orderService.GetAll(page, userID)
	g.JSON(http.StatusOK, page)
}

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

func (c *Controller) CompleteOrder(g *gin.Context) {
	userID := api_helper.GetUserId(g)
	err := c.orderService.CompleteOrder(userID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusCreated, api_helper.Response{Message: "Order Created"})
}

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

func (c *Controller) GetOrders(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	userID := api_helper.GetUserId(g)
	page = c.orderService.GetAll(page, userID)
	g.JSON(http.StatusOK, page)
}
